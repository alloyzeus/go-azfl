package errors

import "strings"

type ContextError interface {
	CallError
	ContextError() ContextError
}

type ContextErrorBuilder interface {
	ContextError

	// Desc returns a copy with descriptor is set to desc.
	Desc(desc ErrorDescriptor) ContextErrorBuilder

	// DescMsg sets the descriptor with the provided string. For the best
	// experience, descMsg should be defined as a constant so that the error
	// users could use it to identify an error. For non-constant descriptor
	// use the Wrap method.
	DescMsg(descMsg string) ContextErrorBuilder

	// Wrap returns a copy with wrapped error is set to detailingError.
	Wrap(detailingError error) ContextErrorBuilder

	Fieldset(fields ...NamedError) ContextErrorBuilder
}

func IsContextError(err error) bool {
	_, ok := err.(ContextError)
	return ok
}

func Context() ContextErrorBuilder {
	return &contextError{}
}

func ContextUnspecified() ContextErrorBuilder {
	return Context().Wrap(ErrValueUnspecified)
}

func IsContextUnspecifiedError(err error) bool {
	if !IsContextError(err) {
		return false
	}
	if desc := UnwrapDescriptor(err); desc != nil {
		return desc == ErrValueUnspecified
	}
	return false
}

type contextError struct {
	descriptor ErrorDescriptor
	wrapped    error
	fields     []NamedError
}

var (
	_ error               = &contextError{}
	_ CallError           = &contextError{}
	_ ContextError        = &contextError{}
	_ ContextErrorBuilder = &contextError{}
	_ Unwrappable         = &contextError{}
)

func (e *contextError) Error() string {
	suffix := e.fieldErrorsAsString()
	if suffix != "" {
		suffix = ": " + suffix
	}
	var descStr string
	if e.descriptor != nil {
		descStr = e.descriptor.Error()
	}
	causeStr := errorString(e.wrapped)
	if causeStr == "" {
		causeStr = descStr
	} else if descStr != "" {
		causeStr = descStr + ": " + causeStr
	}

	if causeStr != "" {
		return "context " + causeStr + suffix
	}
	return "context invalid" + suffix
}

func (e *contextError) CallError() CallError       { return e }
func (e *contextError) ContextError() ContextError { return e }
func (e *contextError) Unwrap() error              { return e.wrapped }

func (e contextError) fieldErrorsAsString() string {
	if flen := len(e.fields); flen > 0 {
		parts := make([]string, 0, flen)
		for _, sub := range e.fields {
			parts = append(parts, sub.Error())
		}
		return strings.Join(parts, ", ")
	}
	return ""
}

func (e contextError) Descriptor() ErrorDescriptor {
	if e.descriptor != nil {
		return e.descriptor
	}
	if desc, ok := e.wrapped.(ErrorDescriptor); ok {
		return desc
	}
	return nil
}

func (e contextError) Desc(desc ErrorDescriptor) ContextErrorBuilder {
	e.descriptor = desc
	return &e
}

func (e contextError) DescMsg(descMsg string) ContextErrorBuilder {
	e.descriptor = constantErrorDescriptor(descMsg)
	return &e
}

func (e contextError) Fieldset(fields ...NamedError) ContextErrorBuilder {
	e.fields = fields // copy?
	return &e
}

func (e contextError) Wrap(detailingError error) ContextErrorBuilder {
	e.wrapped = detailingError
	return &e
}
