package errors

// ValueError describes where the error is.
type ValueError interface {
	error
	ValueError() ValueError
}

type valueConstantErrorDescriptor string

var (
	_ error                = valueConstantErrorDescriptor("")
	_ ErrorDescriptor      = valueConstantErrorDescriptor("")
	_ ValueErrorDescriptor = valueConstantErrorDescriptor("")
)

func (e valueConstantErrorDescriptor) Error() string                      { return string(e) }
func (e valueConstantErrorDescriptor) ErrorDescriptorString() string      { return string(e) }
func (e valueConstantErrorDescriptor) ValueErrorDescriptorString() string { return string(e) }

type ValueErrorDescriptor interface {
	ErrorDescriptor
	ValueErrorDescriptorString() string
}

const (
	// Value was not provided (nil)
	ErrValueUnspecified = valueConstantErrorDescriptor("unspecified")
	// Value was provided but empty
	ErrValueEmpty = valueConstantErrorDescriptor("empty")

	ErrValueMalformed = valueConstantErrorDescriptor("malformed")
	// The value provided does not match the one required by the system
	ErrValueMismatch = valueConstantErrorDescriptor("mismatch")
	// The value provided is currently not supported or not recognized by
	// the system
	ErrValueUnsupported     = valueConstantErrorDescriptor("unsupported")
	ErrValueTypeUnsupported = valueConstantErrorDescriptor("type unsupported")

	// If everything else fails
	ErrValueInvalid = valueConstantErrorDescriptor("invalid")
)

func ValueMalformed(details error) ValueError {
	return &valueDescriptorDetailsError{descriptor: ErrValueMalformed, details: details}
}

func IsValueMalformedError(err error) bool {
	if err == ErrValueMalformed {
		return true
	}
	if desc := UnwrapDescriptor(err); desc != nil {
		return desc == ErrValueMalformed
	}
	return false
}

type valueDescriptorDetailsError struct {
	descriptor valueConstantErrorDescriptor
	details    error
}

var (
	_ error         = valueDescriptorDetailsError{}
	_ ValueError    = valueDescriptorDetailsError{}
	_ hasDescriptor = valueDescriptorDetailsError{}
	_ Unwrappable   = valueDescriptorDetailsError{}
)

func (e valueDescriptorDetailsError) Error() string {
	if e.descriptor != "" {
		if e.details != nil {
			return e.descriptor.Error() + ": " + e.details.Error()
		}
		return e.descriptor.Error()
	}
	if e.details != nil {
		return e.details.Error()
	}
	return ""
}

func (e valueDescriptorDetailsError) Descriptor() ErrorDescriptor { return e.descriptor }
func (e valueDescriptorDetailsError) ValueError() ValueError      { return e }
func (e valueDescriptorDetailsError) Unwrap() error               { return e.details }
