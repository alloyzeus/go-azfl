package errors

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
func Wrap(contextMessage string, causeErr error) error {
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
