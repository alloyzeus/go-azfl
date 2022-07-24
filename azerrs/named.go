package errors

import (
	"fmt"
	"strings"
)

type NamedError interface {
	Unwrappable

	Name() string
	Value() any
}

type NamedErrorBuilder interface {
	NamedError

	// Desc returns a copy with descriptor is set to desc.
	Desc(desc ErrorDescriptor) NamedErrorBuilder

	// DescMsg sets the descriptor with the provided string. For the best
	// experience, descMsg should be defined as a constant so that the error
	// users could use it to identify an error. For non-constant descriptor
	// use the Wrap method.
	DescMsg(descMsg string) NamedErrorBuilder

	// Value provides the value associated to the name.
	Val(value any) NamedErrorBuilder

	// Wrap returns a copy with wrapped error is set to detailingError.
	Wrap(detailingError error) NamedErrorBuilder

	// Rewrap collects descriptor, wrapped, and fields from err and include
	// them into the new error.
	Rewrap(err error) NamedErrorBuilder

	Fieldset(fields ...NamedError) NamedErrorBuilder
}

func N(name string) NamedErrorBuilder {
	return &namedError{name: name}
}

// NamedValueMalformed creates an NamedValue with name is set to
// the value of valueName and descriptor is set to ErrValueMalformed.
func NamedValueMalformed(valueName string) NamedErrorBuilder {
	return &namedError{
		name:       valueName,
		descriptor: ErrValueMalformed,
	}
}

// NamedValueUnsupported creates an NamedValue with name is set to
// the value of valueName and descriptor is set to ErrValueUnsupported.
func NamedValueUnsupported(valueName string) NamedErrorBuilder {
	return &namedError{
		name:       valueName,
		descriptor: ErrValueUnsupported,
	}
}

type namedError struct {
	name       string
	descriptor ErrorDescriptor
	wrapped    error
	value      any
	fields     []NamedError
}

func (e *namedError) Error() string {
	suffix := namedSetToString(e.fields)
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

	var name string
	if e.name != "" {
		if e.value != nil {
			valStr := fmt.Sprintf("%v", e.value)
			if strings.ContainsAny(valStr, string([]byte{'\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0})) {
				name = fmt.Sprintf("%s=%q", e.name, valStr)
			} else {
				name = fmt.Sprintf("%s=%s", e.name, valStr)
			}
		} else {
			name = e.name
		}
	}

	if name != "" {
		if causeStr != "" {
			return name + ": " + causeStr + suffix
		}
		return name + suffix
	}
	if causeStr != "" {
		return causeStr + suffix
	}
	if suffix != "" {
		return strings.TrimPrefix(suffix, ": ")
	}
	return ""
}

func (e namedError) Unwrap() error { return e.wrapped }
func (e namedError) Name() string  { return e.name }
func (e namedError) Value() any    { return e.value }
func (e namedError) Descriptor() ErrorDescriptor {
	if e.descriptor != nil {
		return e.descriptor
	}
	if desc, ok := e.wrapped.(ErrorDescriptor); ok {
		return desc
	}
	return nil
}
func (e namedError) FieldErrors() []NamedError {
	return copyNamedSet(e.fields)
}

func (e namedError) Rewrap(err error) NamedErrorBuilder {
	if err != nil {
		if descErr, _ := err.(ErrorDescriptor); descErr != nil {
			e.descriptor = descErr
			e.value = nil
			e.wrapped = nil
			e.fields = nil
		} else {
			e.descriptor = UnwrapDescriptor(err)
			e.value = unwrapValue(err)
			e.wrapped = Unwrap(err)
			e.fields = UnwrapFieldErrors(err)
		}
	}
	return &e
}

func (e namedError) Desc(desc ErrorDescriptor) NamedErrorBuilder {
	e.descriptor = desc
	return &e
}

func (e namedError) DescMsg(descMsg string) NamedErrorBuilder {
	e.descriptor = constantErrorDescriptor(descMsg)
	return &e
}

func (e namedError) Val(value any) NamedErrorBuilder {
	e.value = value
	return &e
}

func (e namedError) Fieldset(fields ...NamedError) NamedErrorBuilder {
	e.fields = copyNamedSet(fields)
	return &e
}

func (e namedError) Wrap(detailingError error) NamedErrorBuilder {
	e.wrapped = detailingError
	return &e
}

func copyNamedSet(namedSet []NamedError) []NamedError {
	var copiedFields []NamedError
	if len(namedSet) > 0 {
		copiedFields = make([]NamedError, 0, len(namedSet))
		for _, e := range namedSet {
			if e != nil {
				copiedFields = append(copiedFields, e)
			}
		}
	}
	return copiedFields
}

func namedSetToString(namedSet []NamedError) string {
	if flen := len(namedSet); flen > 0 {
		parts := make([]string, 0, flen)
		for _, sub := range namedSet {
			parts = append(parts, sub.Error())
		}
		return strings.Join(parts, ", ")
	}
	return ""
}

func unwrapValue(err error) any {
	if err != nil {
		if d, ok := err.(valueHolder); ok {
			return d.Value()
		}
	}
	return nil
}

type valueHolder interface {
	Value() any
}
