package errors

type hasFieldErrors interface {
	FieldErrors() []NamedError
}

func UnwrapFieldErrors(err error) []NamedError {
	if prov, _ := err.(hasFieldErrors); prov != nil {
		return prov.FieldErrors()
	}
	return nil
}
