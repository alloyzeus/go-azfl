package azcore

// Error is the module error type.
type Error error

// UnwrappableError provides a contract for unwrappable errors.
type UnwrappableError interface {
	Error
	Unwrap() error
}

// CompositeError is an error which able to provide its individual field
// error if any.
type CompositeError interface {
	Error

	FieldErrors() []NamedError
}

// A NamedError describes the error for a composite field.
type NamedError interface {
	UnwrappableError

	// Name returns the name of the subject for this error. This could
	// be argument name or field name.
	Name() string
}
