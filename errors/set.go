package errors

// ErrorSet is an abstraction for an error that holds multiple errors. This
// is usually used for operations that collects all errors before returning.
type ErrorSet interface {
	error
	Errors() []error
}

// UnwrapErrorSet returns the contained error instances if err is
// indeed a ErrorSet.
func UnwrapErrorSet(err error) []error {
	if errSet := asErrorSet(err); errSet != nil {
		return errSet.Errors()
	}
	return nil
}

// asErrorSet returns err as an ErrorSet if err is indeed an ErrorSet,
// otherwise it returns nil.
func asErrorSet(err error) ErrorSet {
	e, _ := err.(ErrorSet)
	return e
}
