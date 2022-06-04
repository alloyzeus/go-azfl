package errors

// ArgumentError abstracts all errors which was caused by error in
// one of the arguments in a function call. This class of error
// has the similar concept as 4xx status codes in HTTP, and thus can
// be mapped into one of these code when used in HTTP request handlers.
type ArgumentError interface {
	CallError
	Unwrappable
	ArgumentName() string
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

func (*argumentError) CallError() {}

func (e *argumentError) Error() string {
	if e.identifier != "" {
		if errMsg := e.innerMsg(); errMsg != "" {
			return "arg " + e.identifier + ": " + errMsg
		}
		return "arg " + e.identifier + " invalid"
	}
	if errMsg := e.innerMsg(); errMsg != "" {
		return "arg " + errMsg
	}
	return "arg invalid"
}

func (e *argumentError) Unwrap() error {
	return e.err
}
