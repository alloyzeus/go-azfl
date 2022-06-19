package azcore

import "github.com/alloyzeus/go-azfl/errors"

// Error is the module error type.
type Error error

const (
	ErrUserContextRequired    = constantErrorDescriptor("user context required")
	ErrServiceContextRequired = constantErrorDescriptor("service context required")
)

func UserContextRequiredError(details error) error {
	return errors.Access(ErrUserContextRequired, details)
}

func ServiceContextRequiredError(details error) error {
	return errors.Access(ErrServiceContextRequired, details)
}

type constantErrorDescriptor string

var (
	_ error = constantErrorDescriptor("")
)

func (e constantErrorDescriptor) Error() string                 { return string(e) }
func (e constantErrorDescriptor) ErrorDescriptorString() string { return string(e) }
