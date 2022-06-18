// Package errors provides extra functionalities to that of Go's stdlib
// errors package.
package errors

import (
	"errors"
)

// Wraps Go's errors
var (
	As     = errors.As
	Is     = errors.Is
	New    = errors.New // Prefer Msg instead as it has better semantic
	Msg    = errors.New
	Unwrap = errors.Unwrap
)

//TODO: fields.
// e.g., errors.Msg("error message", errors.Str("name", name), errors.Err(err))
//    or errors.With().Str("name", name).Err(err).Msg("error message")
//    or errors.With().StrErr("name", nameErr).Msg("error message")
// (sounds like structured logging? exactly!)

// Unwrappable is an error which holds inner error, which usually
// provides the cause if this error.
type Unwrappable interface {
	error
	Unwrap() error
}

// Wrap creates a new error by providing context message to another error.
// It's recommended for the message to describe what the program did which
// caused the error.
//
//     err := fetchData(...)
//     if err != nil { return errors.Wrap("fetching data", err) }
//
func Wrap(contextMessage string, causeErr error) Unwrappable {
	return &errorWrap{contextMessage, causeErr}
}

var _ Unwrappable = &errorWrap{}

type errorWrap struct {
	msg string
	err error
}

func (e *errorWrap) Error() string {
	if e.msg != "" {
		if e.err != nil {
			return e.msg + ": " + e.err.Error()
		}
		return e.msg
	}
	if e.err != nil {
		return e.err.Error()
	}
	return ""
}

func (e *errorWrap) Unwrap() error {
	return e.err
}

const (
	// ErrUnimplemented is used to declare that a functionality, or part of it,
	// has not been implemented. This could be well mapped to some protocols'
	// status code, e.g., HTTP's 501 and gRPC's 12 .
	ErrUnimplemented = errorConstantDescriptor("unimplemented")
)

// A ErrorDescriptor provides the description-part of an error.
//
// In a good practice, an error contains an information about *what* and *why*.
// A ErrorDescriptor abstracts the answers to the *why*.
//
// "User is not found" could be break down into "User" as the answer to the
// *what*, and "not found" as the answer to the *why*. Here, the ErrorDescriptor
// will contain the "not found".
//
// This interface could be used to describe any *what*, like the method
// in "method not implemented". For specific to data, see DataDescriptorError.
type ErrorDescriptor interface {
	error
	ErrorDescriptorString() string
}

type hasDescriptor interface {
	Descriptor() ErrorDescriptor
}

// errorConstantDescriptor is a generic error that designed to be declared
// as constant so that an instance could be easily compared by value to
// its original definiton.
//
// A errorConstantDescriptor is not providing a full context about an error,
// instead it's designed to be wrapped to provide context to the parent error.
type errorConstantDescriptor string

var (
	_ error           = errorConstantDescriptor("")
	_ ErrorDescriptor = errorConstantDescriptor("")
)

func (e errorConstantDescriptor) Error() string                 { return string(e) }
func (e errorConstantDescriptor) ErrorDescriptorString() string { return string(e) }
