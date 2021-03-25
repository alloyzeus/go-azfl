package azcore

import (
	"time"

	"github.com/alloyzeus/go-azfl/azfl/azob"
)

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

//region ServiceMethodCallInfo

// A ServiceMethodCallInfo describes a method call.
type ServiceMethodCallInfo interface {
	AZServiceMethodCallInfo()

	CallID() ServiceMethodCallID

	// RequestOriginTime is the time assigned by the terminal which made the
	// request for the method call.
	//
	// The value is untrusted.
	RequestOriginTime() *time.Time
}

//endregion

//region ServiceMethodCallID

// ServiceMethodCallID represents the identifier of a method call. This
// identifier doubles as idempotency token.
type ServiceMethodCallID interface {
	azob.Equatable

	AZServiceMethodCallID()
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

//region ServiceMethodRequest

// ServiceMethodRequest abstracts method request messages.
type ServiceMethodRequest interface {
	ServiceMethodMessage

	MethodRequestContext() ServiceMethodRequestContext
}

// type ServiceMethodRequestBase struct {
// 	context ServiceMethodRequestContext
// 	parameters ServiceMethodRequestParameters
// }

//endregion

//region ServiceMethodRequestError

// ServiceMethodRequestError is a sub-class of ServiceMethodError which
// indicates that there's an error in the request.
//
// This error class is analogous to HTTP's 4xx status codes.
//
//TODO: sub-classes: acces, parameters, context
type ServiceMethodRequestError interface {
	ServiceMethodError

	AZServiceMethodRequestError()
}

//endregion

//region ServiceMethodRequestContext

// ServiceMethodRequestContext provides an abstraction for all input contexts
// in method call inputs.
type ServiceMethodRequestContext interface {
	ServiceMethodContext

	AZServiceMethodRequestContext()

	// Session returns the session for this context.
	Session() Session
}

//endregion

//region ServiceMethodRequestContextError

// ServiceMethodRequestContextError provides information for
// request-context-related error. It is a sub-class of
// ServiceMethodRequestError.
type ServiceMethodRequestContextError interface {
	ServiceMethodRequestError

	AZServiceMethodRequestContextError()
}

// ServiceMethodRequestSessionError is a sub-class of
// ServiceMethodRequestContextError specialized for indicating error
// in the session.
type ServiceMethodRequestSessionError interface {
	ServiceMethodRequestContextError

	AZServiceMethodRequestSessionError()
}

//endregion

//region ServiceMethodRequestContextBase

// ServiceMethodRequestContextBase is a partial implementation
// of ServiceMethodRequestContext.
type ServiceMethodRequestContextBase struct {
	session Session
}

var _ ServiceMethodContext = ServiceMethodRequestContextBase{}
var _ ServiceMethodRequestContext = ServiceMethodRequestContextBase{}

// AZContext is required
// for conformance with Context.
func (ServiceMethodRequestContextBase) AZContext() {}

// AZServiceContext is required
// for conformance with ServiceContext.
func (ServiceMethodRequestContextBase) AZServiceContext() {}

// AZServiceMethodContext is required
// for conformance with ServiceMethodContext.
func (ServiceMethodRequestContextBase) AZServiceMethodContext() {}

// AZServiceMethodRequestContext is required
// for conformance with ServiceMethodRequestContext.
func (ServiceMethodRequestContextBase) AZServiceMethodRequestContext() {}

// Session is required
// for conformance with ServiceMethodRequestContext.
func (ctx ServiceMethodRequestContextBase) Session() Session { return ctx.session }

//endregion

//region ServiceMethodResponse

// ServiceMethodResponse abstracts method response messages.
type ServiceMethodResponse interface {
	ServiceMethodMessage

	MethodResponseContext() ServiceMethodResponseContext
}

//endregion

//region ServiceMethodResponseContext

// ServiceMethodResponseContext provides an abstraction for all output contexts
// in method call outputs.
//
//TODO: listing of affected states with their respective revision ID.
//TODO: directive: done/end, redirect, retry (on failure; optionally with
// timing and retry count parameters or exponentially back-off parameters)
type ServiceMethodResponseContext interface {
	ServiceMethodContext

	AZServiceMethodResponseContext()

	// // Succeed returns true when the method achieved its objective.
	// Succeed() bool

	// Returns the error, if any.
	Err() ServiceMethodError

	// Mutated returns true if the method made any changes to any state in the
	// server, even when the method did not succeed. It should not
	// return true if the change has been completely rolled-back before the
	// method returned this context.
	Mutated() bool
}

//endregion

//region ServiceMethodResponseContextBase

// ServiceMethodResponseContextBase is a base
// for ServiceMethodResponseContext implementations.
type ServiceMethodResponseContextBase struct {
	err     ServiceMethodError
	mutated bool
}

var _ ServiceMethodResponseContext = ServiceMethodResponseContextBase{}

// NewMethodResponseContext creates a new instance
// of ServiceMethodResponseContext.
func NewMethodResponseContext(
	err ServiceMethodError,
	mutated bool,
) ServiceMethodResponseContextBase {
	return ServiceMethodResponseContextBase{err: err, mutated: mutated}
}

// AZContext is required for conformance with Context.
func (ServiceMethodResponseContextBase) AZContext() {}

// AZServiceContext is required for conformance with ServiceContext.
func (ServiceMethodResponseContextBase) AZServiceContext() {}

// AZServiceMethodContext is required
// for conformance with ServiceMethodContext.
func (ServiceMethodResponseContextBase) AZServiceMethodContext() {}

// AZServiceMethodResponseContext is required
// for conformance with ServiceMethodResponseContext.
func (ServiceMethodResponseContextBase) AZServiceMethodResponseContext() {}

// Err is required for conformance with ServiceMethodResponseContext.
func (ctx ServiceMethodResponseContextBase) Err() ServiceMethodError { return ctx.err }

// Mutated is required for conformance with ServiceMethodResponseContext.
func (ctx ServiceMethodResponseContextBase) Mutated() bool { return ctx.mutated }

//endregion

//region Mutating method

// ServiceMutatingMethodMessage abstracts mutating method requests
// and responses.
type ServiceMutatingMethodMessage interface {
	ServiceMethodMessage

	MutatingMethodContext() ServiceMutatingMethodContext
}

// ServiceMutatingMethodContext abstracts contexts of mutating
// method requests and responses.
type ServiceMutatingMethodContext interface {
	ServiceMethodContext
}

// ServiceMutatingMethodRequest abstracts mutating method requests.
type ServiceMutatingMethodRequest interface {
	ServiceMutatingMethodMessage
	ServiceMethodRequest

	MutatingMethodRequestContext() ServiceMutatingMethodRequestContext
}

// ServiceMutatingMethodRequestContext abstracts mutating method request contexts.
type ServiceMutatingMethodRequestContext interface {
	ServiceMutatingMethodContext
	ServiceMethodRequestContext
}

// ServiceMutatingMethodResponse abstracts mutating method responses.
type ServiceMutatingMethodResponse interface {
	ServiceMutatingMethodMessage
	ServiceMethodResponse

	MutatingMethodResponseContext() ServiceMutatingMethodResponseContext
}

// ServiceMutatingMethodResponseContext abstracts mutating
// method response contexts.
type ServiceMutatingMethodResponseContext interface {
	ServiceMutatingMethodContext
	ServiceMethodResponseContext
}

//endregion
