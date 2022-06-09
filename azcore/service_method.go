package azcore

//region ServiceMethodError

// ServiceMethodError is a specialization of ServiceError which focuses
// on method-related errors.
type ServiceMethodError interface {
	ServiceError

	AZServiceMethodError()
}

// ServiceMethodErrorMsg is a basic implementation of ServiceMethodError
// which provides an error message string.
type ServiceMethodErrorMsg struct {
	msg string
}

var _ ServiceMethodError = ServiceMethodErrorMsg{}

// AZServiceMethodError is required for conformance with ServiceMethodError.
func (ServiceMethodErrorMsg) AZServiceMethodError() {}

// AZServiceError is required for conformance with ServiceError.
func (ServiceMethodErrorMsg) AZServiceError() {}

func (err ServiceMethodErrorMsg) Error() string { return err.msg }

// ServiceMethodInternalError represents error in the method, service, or in
// any dependency required to achieve the objective.
//
// This is analogous to HTTP's 5xx status code.
type ServiceMethodInternalError interface {
	ServiceMethodError

	AZServiceMethodInternalError()
}

// ServiceMethodInternalErrorMsg is a basic implementation of
// ServiceMethodInternalError which provides error message string.
type ServiceMethodInternalErrorMsg struct {
	msg string
}

var _ ServiceMethodInternalError = ServiceMethodInternalErrorMsg{}

// AZServiceMethodInternalError is required for conformance
// with ServiceMethodInternalError.
func (ServiceMethodInternalErrorMsg) AZServiceMethodInternalError() {}

// AZServiceMethodError is required for conformance with ServiceMethodError.
func (ServiceMethodInternalErrorMsg) AZServiceMethodError() {}

// AZServiceError is required for conformance with ServiceError.
func (ServiceMethodInternalErrorMsg) AZServiceError() {}

func (err ServiceMethodInternalErrorMsg) Error() string { return err.msg }

//endregion

// ErrServiceMethodNotImplemented is usually used when a method is unable to
// achieve its objective because some part of it is unimplemented.
//
// Analogous to HTTP's 501 status code and gRPC's 12 status code.
var ErrServiceMethodNotImplemented = &ServiceMethodInternalErrorMsg{msg: "not implemented"}

//region ServiceMethodContext

// ServiceMethodContext is an abstraction for input and output contexts used
// when calling a method.
type ServiceMethodContext interface {
	ServiceContext

	AZServiceMethodContext()
}

//endregion

//region ServiceMethodMessage

// ServiceMethodMessage abstracts the messages, i.e., requests and responses.
type ServiceMethodMessage interface {
	AZServiceMethodMessage()

	// ServiceMethodContext returns the context of this message.
	//
	// Implementations must return most specialized context implementation.
	MethodContext() ServiceMethodContext
}

//endregion
