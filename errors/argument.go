package errors

//TODO: ArgSet

// ArgumentError abstracts all errors which was caused by error in
// one of the arguments in a function call. This class of error
// has the similar concept as 4xx status codes in HTTP, and thus can
// be mapped into one of these code when used in HTTP request handlers.
type ArgumentError interface {
	CallError
	Unwrappable

	// ArgumentName returns the name of the offending argument. It might
	// be empty if there's only one argument.
	ArgumentName() string
}

func IsArgumentError(err error) bool {
	_, ok := err.(ArgumentError)
	return ok
}

func Arg(argName string, err error, fields ...EntityError) ArgumentError {
	return &argumentError{entityError{
		identifier: argName,
		err:        err,
		fields:     fields,
	}}
}

func ArgFields(argName string, fields ...EntityError) ArgumentError {
	return &argumentError{entityError{
		identifier: argName,
		fields:     fields,
	}}
}

func ArgMsg(argName, errMsg string, fields ...EntityError) ArgumentError {
	return &argumentError{entityError{
		identifier: argName,
		err:        Msg(errMsg),
		fields:     fields,
	}}
}

func ArgWrap(argName, contextMessage string, err error, fields ...EntityError) ArgumentError {
	return &argumentError{entityError{
		identifier: argName,
		err:        Wrap(contextMessage, err),
		fields:     fields,
	}}
}

// ArgUnspecified creates an ArgumentError err is set to DataErrUnspecified.
func ArgUnspecified(argName string) ArgumentError {
	return &argumentError{entityError{
		identifier: argName,
		err:        ErrValueUnspecified,
	}}
}

func IsArgUnspecifiedError(err error) bool {
	if !IsArgumentError(err) {
		return false
	}
	if desc := UnwrapDescriptor(err); desc != nil {
		return desc == ErrValueUnspecified
	}
	return false
}

func IsArgUnspecified(err error, argName string) bool {
	if err == nil {
		return false
	}
	argErr, ok := err.(ArgumentError)
	if !ok {
		return false
	}
	if argErr.ArgumentName() != argName {
		return false
	}
	if desc := UnwrapDescriptor(err); desc != nil {
		return desc == ErrValueUnspecified
	}
	return false
}

type argumentError struct {
	entityError
}

var (
	_ error         = &argumentError{}
	_ Unwrappable   = &argumentError{}
	_ CallError     = &argumentError{}
	_ EntityError   = &argumentError{}
	_ ArgumentError = &argumentError{}
)

func (e *argumentError) ArgumentName() string {
	return e.entityError.identifier
}

func (e *argumentError) CallError() CallError { return e }

func (e *argumentError) Error() string {
	suffix := e.fieldErrorsAsString()
	if suffix != "" {
		suffix = ": " + suffix
	}
	detailsStr := e.innerMsg()

	if e.identifier != "" {
		if detailsStr != "" {
			return "arg " + e.identifier + ": " + detailsStr + suffix
		}
		return "arg " + e.identifier + suffix
	}
	if detailsStr != "" {
		return "arg " + detailsStr + suffix
	}
	if suffix != "" {
		return "arg" + suffix
	}
	return "arg error"
}

func (e *argumentError) Unwrap() error {
	return e.err
}
