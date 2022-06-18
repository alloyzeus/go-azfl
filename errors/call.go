package errors

import "strings"

type CallError interface {
	error
	CallError() CallError
}

func IsCallError(e error) bool {
	_, ok := e.(CallError)
	return ok
}

// CallErrorSet is a facility to combine multiple CallError instances into a
// single CallError.
//
// In some cases, it's prefered to have all call-related errors for a call
// a list. This way, the user could fix them all in one go.
type CallErrorSet interface {
	CallError
	CallErrorSet() []CallError
}

type callErrorSet struct {
	callErrors []CallError
}

var (
	_ error = &callErrorSet{}
)

func (e *callErrorSet) Error() string {
	if len(e.callErrors) > 0 {
		errs := make([]string, 0, len(e.callErrors))
		for _, ce := range e.callErrors {
			errs = append(errs, ce.Error())
		}
		s := strings.Join(errs, ", ")
		if s != "" {
			return "call: " + s
		}
	}
	return "call error"
}
