package errors

// A ErrorDescriptor provides the description-part of an error.
//
// In a good practice, an error contains an information about *what* and *why*.
// A ErrorDescriptor abstracts the answers to the *why*.
//
// "User is not found" could be break down into "User" as the answer to the
// *what*, and "not found" as the answer to the *why*. Here, the ErrorDescriptor
// will contain the "not found".
//
// This interface could be used to describe any *what*, like the method
// in "method not implemented". For specific to data, see DataDescriptorError.
type ErrorDescriptor interface {
	error
	ErrorDescriptorString() string
}

type hasDescriptor interface {
	Descriptor() ErrorDescriptor
}

func DescWrap(descriptor ErrorDescriptor, details error) error {
	return descriptorDetailsError{descriptor: descriptor, details: details}
}

type descriptorDetailsError struct {
	descriptor ErrorDescriptor
	details    error
}

var (
	_ error         = descriptorDetailsError{}
	_ hasDescriptor = descriptorDetailsError{}
	_ Unwrappable   = descriptorDetailsError{}
)

func (e descriptorDetailsError) Error() string {
	if e.descriptor != nil {
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
func (e descriptorDetailsError) Unwrap() error               { return e.details }

// constantErrorDescriptor is a generic error that designed to be declared
// as constant so that an instance could be easily compared by value to
// its original definiton.
//
// A constantErrorDescriptor is not providing a full context about an error,
// instead it's designed to be wrapped to provide context to the parent error.
type constantErrorDescriptor string

var (
	_ error           = constantErrorDescriptor("")
	_ ErrorDescriptor = constantErrorDescriptor("")
)

func (e constantErrorDescriptor) Error() string                 { return string(e) }
func (e constantErrorDescriptor) ErrorDescriptorString() string { return string(e) }

func UnwrapDescriptor(err error) ErrorDescriptor {
	if err != nil {
		if d, ok := err.(hasDescriptor); ok {
			return d.Descriptor()
		}
	}
	return nil
}

func HasDescriptor(err error, desc ErrorDescriptor) bool {
	d := UnwrapDescriptor(err)
	//TODO: use other strategies. even reflect as the last resort
	return d == desc
}

func HasDescriptorText(err error, descText string) bool {
	d := UnwrapDescriptor(err)
	return d.Error() == descText
}

func errorDescriptorString(err error) string {
	if err != nil {
		if desc, ok := err.(ErrorDescriptor); ok {
			return desc.ErrorDescriptorString()
		}
	}
	return ""
}
