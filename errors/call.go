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

// CallErrorSet is an interface to combine multiple CallError instances into a
// single CallError.
//
// In some cases, it's prefered to have all call-related errors for a call
// a list. This way, the user could fix them all in one go.
type CallErrorSet interface {
	CallError
	ErrorSet
	CallErrors() []CallError
}

// UnwrapCallErrorSet returns the contained CallError instances if err is
// indeed a CallErrorSet.
func UnwrapCallErrorSet(err error) []CallError {
	if errSet := asCallErrorSet(err); errSet != nil {
		return errSet.CallErrors()
	}
	return nil
}

// asCallErrorSet returrns err as CallErrorSet if err is indeed a CallErrorSet,
// otherwise it returrns nil.
func asCallErrorSet(err error) CallErrorSet {
	e, _ := err.(CallErrorSet)
	return e
}

// CallSet creates a compound error comprised of multiple instances of
// CallError. The resulting error is also a CallError.
func CallSet(callErrors ...CallError) CallErrorSet {
	ours := make([]CallError, len(callErrors))
	copy(ours, callErrors)
	return callErrorSet(ours)
}

type callErrorSet []CallError

var (
	_ error        = callErrorSet{}
	_ CallError    = callErrorSet{}
	_ callErrorSet = callErrorSet{}
	_ ErrorSet     = callErrorSet{}
)

func (e callErrorSet) Error() string {
	if len(e) > 0 {
		errs := make([]string, 0, len(e))
		for _, ce := range e {
			errs = append(errs, ce.Error())
		}
		s := strings.Join(errs, ", ")
		if s != "" {
			return "call: " + s
		}
	}
	return "call error"
}

func (e callErrorSet) CallErrors() []CallError {
	if len(e) > 0 {
		errs := make([]CallError, len(e))
		copy(errs, e)
		return errs
	}
	return nil
}

func (e callErrorSet) Errors() []error {
	if len(e) > 0 {
		errs := make([]error, 0, len(e))
		for _, i := range e {
			errs = append(errs, i)
		}
		return errs
	}
	return nil
}

func (e callErrorSet) CallError() CallError { return e }
