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

const (
	// ErrUnimplemented is used to declare that a functionality, or part of it,
	// has not been implemented. This could be well mapped to some protocols'
	// status code, e.g., HTTP's 501 and gRPC's 12 .
	ErrUnimplemented = constantErrorDescriptor("unimplemented")
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

// constantErrorDescriptor is a generic error that designed to be declared
// as constant so that an instance could be easily compared by value to
// its original definiton.
//
// A constantErrorDescriptor is not providing a full context about an error,
// instead it's designed to be wrapped to provide context to the parent error.
type constantErrorDescriptor string

var (
	_ error           = constantErrorDescriptor("")
	_ ErrorDescriptor = constantErrorDescriptor("")
)

func (e constantErrorDescriptor) Error() string                 { return string(e) }
func (e constantErrorDescriptor) ErrorDescriptorString() string { return string(e) }

func UnwrapDescriptor(err error) ErrorDescriptor {
	if err != nil {
		if d, ok := err.(hasDescriptor); ok {
			return d.Descriptor()
		}
	}
	return nil
}

func errorDescriptorString(err error) string {
	if err != nil {
		if desc, ok := err.(ErrorDescriptor); ok {
			return desc.ErrorDescriptorString()
		}
	}
	return ""
}

func errorString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
