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
	DataErrUnspecified = dataErrorConstantDescriptor("unspecified")
	// Data was provided but empty
	DataErrEmpty = dataErrorConstantDescriptor("empty")

	DataErrInvalid         = dataErrorConstantDescriptor("invalid")
	DataErrMalformed       = dataErrorConstantDescriptor("malformed")
	DataErrTypeUnsupported = dataErrorConstantDescriptor("type unsupported")
)

func DataMalformed(details error) DataError {
	return &descriptorDetailsError{descriptor: DataErrMalformed, details: details}
}

func IsDataMalformed(err error) bool {
	if err == DataErrMalformed {
		return true
	}
	if d, ok := err.(hasDescriptor); ok {
		desc := d.Descriptor()
		if desc == DataErrMalformed {
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
	return e.details.Error()
}

func (e descriptorDetailsError) Descriptor() ErrorDescriptor { return e.descriptor }
func (e descriptorDetailsError) DataError() DataError        { return e }
func (e descriptorDetailsError) Unwrap() error               { return e.details }
