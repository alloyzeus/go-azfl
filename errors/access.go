package errors

// An AccessError is a specialization of ContextError.
type AccessError interface {
	ContextError
	AccessError() AccessError
}

func IsAccessError(err error) bool {
	_, ok := err.(AccessError)
	return ok
}

func Access(descriptor ErrorDescriptor, details error) AccessError {
	return &accessError{
		descriptor: descriptor,
		details:    details,
	}
}

const ErrAccessForbidden = accessErrorDescriptor("forbidden")

func AccessForbidden(details error) AccessError {
	return &accessError{
		descriptor: ErrAccessForbidden,
		details:    details,
	}
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

func AuthorizationInvalid(details error) AccessError {
	return &accessError{
		descriptor: ErrAuthorizationInvalid,
		details:    details,
	}
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
	details    error
}

var (
	_ error         = &accessError{}
	_ Unwrappable   = &accessError{}
	_ AccessError   = &accessError{}
	_ hasDescriptor = &accessError{}
)

func (e *accessError) Error() string {
	detailsStr := errorString(e.details)
	if descStr := errorDescriptorString(e.descriptor); descStr != "" {
		if detailsStr != "" {
			return "access " + descStr + ": " + detailsStr
		}
		return "access " + descStr
	}
	if detailsStr != "" {
		return "access: " + detailsStr
	}
	return "access error"
}

func (e *accessError) AccessError() AccessError    { return e }
func (e *accessError) ContextError() ContextError  { return e }
func (e *accessError) CallError() CallError        { return e }
func (e *accessError) Descriptor() ErrorDescriptor { return e.descriptor }
func (e *accessError) Unwrap() error               { return e.details }

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