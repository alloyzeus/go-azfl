package errors

// An AccessError is a specialization of ContextError.
type AccessError interface {
	ContextError
	AccessError() AccessError
}

type AccessErrorBuilder interface {
	AccessError

	// Desc returns a copy with descriptor is set to desc.
	Desc(desc ErrorDescriptor) AccessErrorBuilder

	// DescMsg sets the descriptor with the provided string. For the best
	// experience, descMsg should be defined as a constant so that the error
	// users could use it to identify an error. For non-constant descriptor
	// use the Wrap method.
	DescMsg(descMsg string) AccessErrorBuilder

	// Wrap returns a copy with wrapped error is set to detailingError.
	Wrap(detailingError error) AccessErrorBuilder
}

func IsAccessError(err error) bool {
	_, ok := err.(AccessError)
	return ok
}

//TODO: should be Access(resourceIdentifiers ...string)
func Access() AccessErrorBuilder {
	return &accessError{}
}

const ErrAccessForbidden = accessErrorDescriptor("forbidden")

func AccessForbidden() AccessErrorBuilder {
	return Access().Desc(ErrAccessForbidden)
}

func IsAccessForbidden(err error) bool {
	if err == ErrAccessForbidden {
		return true
	}
	if desc := UnwrapDescriptor(err); desc != nil {
		return desc == ErrAccessForbidden
	}
	return false
}

const ErrAuthorizationInvalid = constantErrorDescriptor("authorization invalid")

func AuthorizationInvalid() AccessErrorBuilder {
	return Access().Desc(ErrAuthorizationInvalid)
}

func IsAuthorizationInvalid(err error) bool {
	if err == ErrAuthorizationInvalid {
		return true
	}
	if desc := UnwrapDescriptor(err); desc != nil {
		return desc == ErrAuthorizationInvalid
	}
	return false
}

type accessError struct {
	descriptor ErrorDescriptor
	wrapped    error
}

var (
	_ error         = &accessError{}
	_ Unwrappable   = &accessError{}
	_ AccessError   = &accessError{}
	_ hasDescriptor = &accessError{}
)

func (e *accessError) Error() string {
	causeStr := errorString(e.wrapped)
	if descStr := errorDescriptorString(e.descriptor); descStr != "" {
		if causeStr != "" {
			return "access " + descStr + ": " + causeStr
		}
		return "access " + descStr
	}
	if causeStr != "" {
		return "access: " + causeStr
	}
	return "access error"
}

func (e *accessError) AccessError() AccessError    { return e }
func (e *accessError) ContextError() ContextError  { return e }
func (e *accessError) CallError() CallError        { return e }
func (e *accessError) Descriptor() ErrorDescriptor { return e.descriptor }
func (e *accessError) Unwrap() error               { return e.wrapped }

func (e accessError) Desc(desc ErrorDescriptor) AccessErrorBuilder {
	e.descriptor = desc
	return &e
}

func (e accessError) DescMsg(descMsg string) AccessErrorBuilder {
	e.descriptor = constantErrorDescriptor(descMsg)
	return &e
}

func (e accessError) Wrap(detailingError error) AccessErrorBuilder {
	e.wrapped = detailingError
	return &e
}

type accessErrorDescriptor string

var (
	_ error           = accessErrorDescriptor("")
	_ AccessError     = accessErrorDescriptor("")
	_ ErrorDescriptor = accessErrorDescriptor("")
)

func (e accessErrorDescriptor) AccessError() AccessError      { return e }
func (e accessErrorDescriptor) CallError() CallError          { return e }
func (e accessErrorDescriptor) ContextError() ContextError    { return e }
func (e accessErrorDescriptor) Error() string                 { return string(e) }
func (e accessErrorDescriptor) ErrorDescriptorString() string { return string(e) }
