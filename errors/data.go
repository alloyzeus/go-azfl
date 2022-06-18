package errors

// DataError describes where the error is.
type DataError interface {
	error
	DataError() DataError
}

type dataErrorConstantDescriptor string

var (
	_ error               = dataErrorConstantDescriptor("")
	_ ErrorDescriptor     = dataErrorConstantDescriptor("")
	_ DataErrorDescriptor = dataErrorConstantDescriptor("")
)

func (e dataErrorConstantDescriptor) Error() string                     { return string(e) }
func (e dataErrorConstantDescriptor) ErrorDescriptorString() string     { return string(e) }
func (e dataErrorConstantDescriptor) DataErrorDescriptorString() string { return string(e) }

type DataErrorDescriptor interface {
	ErrorDescriptor
	DataErrorDescriptorString() string
}

const (
	// Data was not provided (nil)
	ErrDataUnspecified = dataErrorConstantDescriptor("unspecified")
	// Data was provided but empty
	ErrDataEmpty = dataErrorConstantDescriptor("empty")

	ErrDataInvalid         = dataErrorConstantDescriptor("invalid")
	ErrDataMalformed       = dataErrorConstantDescriptor("malformed")
	ErrDataTypeUnsupported = dataErrorConstantDescriptor("type unsupported")
)

func DataMalformed(details error) DataError {
	return &descriptorDetailsError{descriptor: ErrDataMalformed, details: details}
}

func IsDataMalformedError(err error) bool {
	if err == ErrDataMalformed {
		return true
	}
	if d, ok := err.(hasDescriptor); ok {
		desc := d.Descriptor()
		if desc == ErrDataMalformed {
			return true
		}
	}
	return false
}

type descriptorDetailsError struct {
	descriptor dataErrorConstantDescriptor
	details    error
}

var (
	_ error         = descriptorDetailsError{}
	_ DataError     = descriptorDetailsError{}
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
func (e descriptorDetailsError) DataError() DataError        { return e }
func (e descriptorDetailsError) Unwrap() error               { return e.details }
