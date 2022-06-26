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

func errorString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
