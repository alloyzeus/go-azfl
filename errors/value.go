package errors

// ValueError describes where the error is.
type ValueError interface {
	error
	ValueError() ValueError
}

type valueErrorConstantDescriptor string

var (
	_ error                = valueErrorConstantDescriptor("")
	_ ErrorDescriptor      = valueErrorConstantDescriptor("")
	_ ValueErrorDescriptor = valueErrorConstantDescriptor("")
)

func (e valueErrorConstantDescriptor) Error() string                      { return string(e) }
func (e valueErrorConstantDescriptor) ErrorDescriptorString() string      { return string(e) }
func (e valueErrorConstantDescriptor) ValueErrorDescriptorString() string { return string(e) }

type ValueErrorDescriptor interface {
	ErrorDescriptor
	ValueErrorDescriptorString() string
}

const (
	// Value was not provided (nil)
	ErrValueUnspecified = valueErrorConstantDescriptor("unspecified")
	// Value was provided but empty
	ErrValueEmpty = valueErrorConstantDescriptor("empty")

	ErrValueInvalid         = valueErrorConstantDescriptor("invalid")
	ErrValueMalformed       = valueErrorConstantDescriptor("malformed")
	ErrValueTypeUnsupported = valueErrorConstantDescriptor("type unsupported")
)

func ValueMalformed(details error) ValueError {
	return &descriptorDetailsError{descriptor: ErrValueMalformed, details: details}
}

func IsValueMalformedError(err error) bool {
	if err == ErrValueMalformed {
		return true
	}
	if d, ok := err.(hasDescriptor); ok {
		desc := d.Descriptor()
		if desc == ErrValueMalformed {
			return true
		}
	}
	return false
}

type descriptorDetailsError struct {
	descriptor valueErrorConstantDescriptor
	details    error
}

var (
	_ error         = descriptorDetailsError{}
	_ ValueError    = descriptorDetailsError{}
	_ hasDescriptor = descriptorDetailsError{}
	_ Unwrappable   = descriptorDetailsError{}
)

func (e descriptorDetailsError) Error() string {
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

func (e descriptorDetailsError) Descriptor() ErrorDescriptor { return e.descriptor }
func (e descriptorDetailsError) ValueError() ValueError      { return e }
func (e descriptorDetailsError) Unwrap() error               { return e.details }
