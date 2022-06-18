package errors

import "strings"

type ContextError interface {
	error
	CallError
	ContextError() ContextError
}

func IsContextError(err error) bool {
	_, ok := err.(ContextError)
	return ok
}

func Context(details error, fields ...EntityError) ContextError {
	return &contextError{inner: details, fields: fields[:]}
}

func ContextFields(fields ...EntityError) ContextError {
	return &contextError{
		fields: fields[:],
	}
}

func ContextUnspecified() ContextError {
	return &contextError{inner: ErrValueUnspecified}
}

func IsContextUnspecifiedError(err error) bool {
	if !IsContextError(err) {
		return false
	}
	if d, ok := err.(hasDescriptor); ok {
		desc := d.Descriptor()
		if desc == ErrValueUnspecified {
			return true
		}
	}
	return false
}

type contextError struct {
	inner  error
	fields []EntityError
}

var (
	_ error        = &contextError{}
	_ CallError    = &contextError{}
	_ ContextError = &contextError{}
	_ Unwrappable  = &contextError{}
)

func (e *contextError) Error() string {
	suffix := e.fieldErrorsAsString()
	if suffix != "" {
		suffix = ": " + suffix
	}

	if errMsg := e.innerMsg(); errMsg != "" {
		return "context " + errMsg + suffix
	}
	return "context invalid" + suffix
}

func (e *contextError) CallError() CallError       { return e }
func (e *contextError) ContextError() ContextError { return e }
func (e *contextError) Unwrap() error              { return e.inner }

func (e *contextError) innerMsg() string {
	if e.inner != nil {
		return e.inner.Error()
	}
	return ""
}

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

func (e *contextError) Descriptor() ErrorDescriptor {
	if e == nil {
		return nil
	}
	if desc, ok := e.inner.(ErrorDescriptor); ok {
		return desc
	}
	if d, ok := e.inner.(hasDescriptor); ok {
		return d.Descriptor()
	}
	return nil
}
