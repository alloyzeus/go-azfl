package errors

// A ErrorDescriptor provides the description-part of an error. It was
// designed to be translated to error codes.
//
// In a good practice, an error contains an information about *what* and *why*.
// A ErrorDescriptor abstracts the answers to the *why*.
//
// "User is not found" could be break down into "User" as the answer to the
// *what*, and "not found" as the answer to the *why*. Here, the ErrorDescriptor
// will contain the "not found".
//
// This interface could be used to describe any *what*, like the method
// in "method not implemented". For specific to value, see ValueDescriptorError.
type ErrorDescriptor interface {
	error
	ErrorDescriptorString() string
}

type DescriptorWrappedError interface {
	Unwrappable

	Descriptor() ErrorDescriptor
}

type hasDescriptor interface {
	Descriptor() ErrorDescriptor
}

func descWrap(descriptor ErrorDescriptor, details error) DescriptorWrappedError {
	return descriptorWrappedError{descriptor: descriptor, wrapped: details}
}

type descriptorWrappedError struct {
	descriptor ErrorDescriptor
	wrapped    error
}

var (
	_ error         = descriptorWrappedError{}
	_ hasDescriptor = descriptorWrappedError{}
	_ Unwrappable   = descriptorWrappedError{}
)

func (e descriptorWrappedError) Error() string {
	if e.descriptor != nil {
		if e.wrapped != nil {
			return e.descriptor.Error() + ": " + e.wrapped.Error()
		}
		return e.descriptor.Error()
	}
	if e.wrapped != nil {
		return e.wrapped.Error()
	}
	return ""
}

func (e descriptorWrappedError) Descriptor() ErrorDescriptor { return e.descriptor }
func (e descriptorWrappedError) Unwrap() error               { return e.wrapped }

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
