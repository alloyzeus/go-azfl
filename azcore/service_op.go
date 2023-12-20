package azcore

//region ServiceOpError

// ServiceOpError is a specialization of ServiceError which focuses
// on method-related errors.
type ServiceOpError interface {
	ServiceError

	AZServiceOpError()
}

// ServiceOpErrorMsg is a basic implementation of ServiceOpError
// which provides an error message string.
type ServiceOpErrorMsg struct {
	msg string
}

var _ ServiceOpError = ServiceOpErrorMsg{}

// AZServiceOpError is required for conformance with ServiceOpError.
func (ServiceOpErrorMsg) AZServiceOpError() {}

// AZServiceError is required for conformance with ServiceError.
func (ServiceOpErrorMsg) AZServiceError() {}

func (err ServiceOpErrorMsg) Error() string { return err.msg }

// ServiceOpInternalError represents error in the method, service, or in
// any dependency required to achieve the objective.
//
// This is analogous to HTTP's 5xx status code.
type ServiceOpInternalError interface {
	ServiceOpError

	AZServiceOpInternalError()
}

// ServiceOpInternalErrorMsg is a basic implementation of
// ServiceOpInternalError which provides error message string.
type ServiceOpInternalErrorMsg struct {
	msg string
}

var _ ServiceOpInternalError = ServiceOpInternalErrorMsg{}

// AZServiceOpInternalError is required for conformance
// with ServiceOpInternalError.
func (ServiceOpInternalErrorMsg) AZServiceOpInternalError() {}

// AZServiceOpError is required for conformance with ServiceOpError.
func (ServiceOpInternalErrorMsg) AZServiceOpError() {}

// AZServiceError is required for conformance with ServiceError.
func (ServiceOpInternalErrorMsg) AZServiceError() {}

func (err ServiceOpInternalErrorMsg) Error() string { return err.msg }

//endregion

// ErrServiceOpNotImplemented is usually used when a method is unable to
// achieve its objective because some part of it is unimplemented.
//
// Analogous to HTTP's 501 status code and gRPC's 12 status code.
var ErrServiceOpNotImplemented = &ServiceOpInternalErrorMsg{msg: "not implemented"}

//region ServiceOpContext

// ServiceOpContext is an abstraction for input and output contexts used
// when calling a method.
type ServiceOpContext interface {
	ServiceContext

	// OperationName returns the name of the method or the endpoint.
	//
	// For HTTP, this method returns the method. For other protocols, it should
	// be the name of the method e.g., `getUser`.
	OperationName() string

	// ResourceID returns the identifier of the resource being accessed by
	// the call. For HTTP, it's the path. For other protocols, it should
	// be the identifier (ID) of the entity. If there's more than one
	// resources, e.g., a relationship between two entities, then it returns
	// the identifiers of the entities separated by commas.
	ResourceID() string
}

//endregion

//region ServiceOpMessage

// ServiceOpMessage abstracts the messages, i.e., requests and responses.
type ServiceOpMessage interface {
	AZServiceOpMessage()

	// ServiceOpContext returns the context of this message.
	//
	// Implementations must return most specialized context implementation.
	ServiceOpContext() ServiceOpContext
}

//endregion
