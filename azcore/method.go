package azcore

import "time"

//region MethodError

// MethodError abstracts all method-related errors.
type MethodError interface {
	Error

	AZMethodError()
}

// MethodErrorMsg is a basic implementation of MethodError which provides
// an error message string.
type MethodErrorMsg struct {
	msg string
}

var _ MethodError = MethodErrorMsg{}

// AZMethodError is required for conformance with MethodError.
func (MethodErrorMsg) AZMethodError() {}

func (err MethodErrorMsg) Error() string { return err.msg }

// MethodInternalError represents error in the method, service, or in
// any dependency required to achieve the objective.
//
// This is analogous to HTTP's 5xx status code.
type MethodInternalError interface {
	MethodError

	AZMethodInternalError()
}

// MethodInternalErrorMsg is a basic implementation of MethodInternalError
// which provides error message string.
type MethodInternalErrorMsg struct {
	msg string
}

var _ MethodInternalError = MethodInternalErrorMsg{}

// AZMethodInternalError is required for conformance with MethodInternalError.
func (MethodInternalErrorMsg) AZMethodInternalError() {}

// AZMethodError is required for conformance with MethodError.
func (MethodInternalErrorMsg) AZMethodError() {}

func (err MethodInternalErrorMsg) Error() string { return err.msg }

//endregion

// ErrMethodNotImplemented is usually used when a method is unable to
// achieve its objective because some part of it is unimplemented.
//
// Analogous to HTTP's 501 status code and gRPC's 12 status code.
var ErrMethodNotImplemented = &MethodInternalErrorMsg{msg: "not implemented"}

//region MethodContext

// MethodContext is an abstraction for input and output contexts used
// when calling a method.
type MethodContext interface {
	Context

	AZMethodContext()
}

//endregion

//region MethodCallInfo

// A MethodCallInfo describes a method call.
type MethodCallInfo interface {
	AZMethodCallInfo()

	CallID() MethodCallID

	// RequestTime is the time assigned by the terminal which made the
	// request for the method call.
	//
	// The value is untrusted.
	RequestTime() *time.Time
}

//endregion

//region MethodCallID

// MethodCallID represents the identifier of a method call. This identifier
// doubles as idempotency token.
type MethodCallID interface {
	Equatable

	AZMethodCallID()
}

//endregion

//region MethodMessage

// MethodMessage abstracts the messages, i.e., requests and responses.
type MethodMessage interface {
	AZMethodMessage()

	// MethodContext returns the context of this message.
	//
	// Implementations must return most specialized context implementation.
	MethodContext() MethodContext
}

//endregion

//region MethodRequest

// MethodRequest abstracts method request messages.
type MethodRequest interface {
	MethodMessage

	MethodRequestContext() MethodRequestContext
}

// type MethodRequestBase struct {
// 	context MethodRequestContext
// 	parameters MethodRequestParameters
// }

//endregion

//region MethodRequestError

// MethodRequestError is a sub-class of MethodError which indicates that
// there's an error in the request.
//
// This error class is analogous to HTTP's 4xx status codes.
//
//TODO: sub-classes: acces, parameters, context
type MethodRequestError interface {
	MethodError

	AZMethodRequestError()
}

//endregion

//region MethodRequestContext

// MethodRequestContext provides an abstraction for all input contexts
// in method call inputs.
type MethodRequestContext interface {
	MethodContext

	AZMethodRequestContext()
}

//endregion

//region MethodRequestContextError

// MethodRequestContextError provides request-context-related error
// information. It is a sub-class of MethodRequestError.
type MethodRequestContextError interface {
	MethodRequestError

	AZMethodRequestContextError()
}

//endregion

//region MethodRequestContextBase

// MethodRequestContextBase is a partial implementation
// of MethodRequestContext.
type MethodRequestContextBase struct{}

var _ MethodContext = MethodRequestContextBase{}
var _ MethodRequestContext = MethodRequestContextBase{}

// AZContext is required
// for conformance with Context.
func (MethodRequestContextBase) AZContext() {}

// AZMethodContext is required
// for conformance with MethodContext.
func (MethodRequestContextBase) AZMethodContext() {}

// AZMethodRequestContext is required
// for conformance with MethodRequestContext.
func (MethodRequestContextBase) AZMethodRequestContext() {}

//endregion

//region MethodResponse

// MethodResponse abstracts method response messages.
type MethodResponse interface {
	MethodMessage

	MethodResponseContext() MethodResponseContext
}

//endregion

//region MethodResponseContext

// MethodResponseContext provides an abstraction for all output contexts
// in method call outputs.
//
//TODO: listing of affected states with their respective revision ID.
//TODO: directive: done/end, retry (on failure; optionally with timing and
// retry count parameters or exponentially back-off parameters), redirect
type MethodResponseContext interface {
	MethodContext

	AZMethodResponseContext()

	// // Succeed returns true when the method achieved its objective.
	// Succeed() bool

	// Returns the error, if any.
	Err() MethodError

	// Mutated returns true if the method made any changes to any state in the
	// server, even when the method did not succeed. It should not
	// return true if the change has been completely rolled-back before the
	// method returned this context.
	Mutated() bool
}

//endregion

//region MethodResponseContextBase

// MethodResponseContextBase is a base
// for MethodResponseContext implementations.
type MethodResponseContextBase struct {
	err     MethodError
	mutated bool
}

var _ MethodResponseContext = MethodResponseContextBase{}

// NewMethodResponseContext creates a new instance
// of MethodResponseContext.
func NewMethodResponseContext(
	err MethodError, mutated bool,
) MethodResponseContextBase {
	return MethodResponseContextBase{err: err, mutated: mutated}
}

// AZContext is required for conformance with Context.
func (MethodResponseContextBase) AZContext() {}

// AZMethodContext is required
// for conformance with MethodContext.
func (MethodResponseContextBase) AZMethodContext() {}

// AZMethodResponseContext is required
// for conformance with MethodResponseContext.
func (MethodResponseContextBase) AZMethodResponseContext() {}

// Err is required for conformance with MethodResponseContext.
func (ctx MethodResponseContextBase) Err() MethodError { return ctx.err }

// Mutated is required for conformance with MethodResponseContext.
func (ctx MethodResponseContextBase) Mutated() bool { return ctx.mutated }

//endregion

//region Mutating method

// MutatingMessage abstracts mutating method requests and responses.
type MutatingMessage interface {
	MethodMessage

	MutatingContext() MutatingContext
}

// MutatingContext abstracts contexts of mutating method requests and responses.
type MutatingContext interface {
	MethodContext
}

// MutatingRequest abstracts mutating method requests.
type MutatingRequest interface {
	MutatingMessage
	MethodRequest

	MutatingRequestContext() MutatingRequestContext
}

// MutatingRequestContext abstracts mutating method request contexts.
type MutatingRequestContext interface {
	MutatingContext
	MethodRequestContext
}

// MutatingResponse abstracts mutating method responses.
type MutatingResponse interface {
	MutatingMessage
	MethodResponse

	MutatingResponseContext() MutatingResponseContext
}

// MutatingResponseContext abstracts mutating method response contexts.
type MutatingResponseContext interface {
	MutatingContext
	MethodResponseContext
}

//endregion
