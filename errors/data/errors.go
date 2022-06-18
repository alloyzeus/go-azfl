// Package data is deprecated. Use the ones from parent package instead.
package data

// Error is an abstract error type for all data-related errors.
type Error interface {
	error
	ValueError() Error
}

func Err(err error) Error { return &wrappingError{err} }

type wrappingError struct {
	err error
}

var _ Error = &wrappingError{}

func (e *wrappingError) Error() string     { return e.err.Error() }
func (e *wrappingError) ValueError() Error { return e }

type msgError struct {
	msg string
}

var _ Error = &msgError{}

func (e *msgError) Error() string     { return e.msg }
func (e *msgError) ValueError() Error { return e }

func Malformed(err error) Error {
	return &malformedError{err}
}

type malformedError struct {
	err error
}

var _ Error = &malformedError{}

func (e *malformedError) Error() string {
	if e.err != nil {
		return "malformed: " + e.err.Error()
	}
	return "malformed"
}

func (e *malformedError) ValueError() Error { return e }

var (
	ErrEmpty           = &msgError{"empty"}
	ErrMalformed       = Malformed(nil)
	ErrTypeUnsupported = &msgError{"type unsupported"}
)
