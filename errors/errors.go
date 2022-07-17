// Package errors provides extra functionalities to that of Go's stdlib
// errors package.
//
// An error as follow is informative for human developers:
//
//     Unable to parse "example" as an email address.
//
// But for other part of code that calls the function, it will need to
// parse that message so that it can react properly, including to translate
// the message into various human languages.
//
// It is pretty much clear that structured errors are more flexible. For the
// calling routines, they won't need to parse the message. If we are going
// to display the error to human end users, the structured errors will also
// have more advantages as different human languages are constructed
// differently.
//
// Examples:
//
//     // An error that describes that argument "username" is unspecified
//     err := errors.Arg("username").Unspecified()
//
//     // Check if an error is an argument error where "username" is unspecified
//     errors.IsArgUnspecified(err, "username")
//
//     // ... or simply check if it's an argument error
//     if errors.IsArgumentError(err) {
//         // e.g., respond with HTTP 400
//     }
//
//     // An error that describes that the field "Name" of entity with ID "user5432" is empty
//     errors.Ent("user5432").Fieldset(errors.Ent("Name").Desc(errors.ErrValueEmpty))
//
//     // An error that describes that argument "email" is malformed. Detailing error, usually from
//     // the parser function, is available as wrapped error.
//     errors.Arg("email").Desc(errors.ErrValueMalformed).Wrap(err)
//
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
