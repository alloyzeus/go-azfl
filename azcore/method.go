package azcore

//region MethodCallContext

// MethodCallContext is an abstraction for input and output contexts used
// when calling a method.
type MethodCallContext interface {
	Context

	AZMethodCallContext()
}

//endregion

//region MethodCallError

// MethodCallError provides an abstraction for all errors returned by a method.
//
//TODO: sub-classes: input (acces, parameters, context), internal
type MethodCallError interface {
	Error

	AZMethodCallError()
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

//region MethodCallMessage

// MethodCallMessage abstracts the messages, i.e., request and response.
type MethodCallMessage interface {
	AZMethodCallMessage()
}

//endregion

//region MethodRequest

// MethodRequest abstracts request messages.
type MethodRequest interface {
	MethodCallMessage

	Context() MethodRequestContext
}

// type MethodRequestBase struct {
// 	context MethodRequestContext
// 	parameters MethodRequestParameters
// }

//endregion

//region MethodRequestError

// MethodRequestError is a sub-class of MethodCallError which indicates that
// there's an error in the request.
type MethodRequestError interface {
	MethodCallError

	AZMethodRequestError()
}

//endregion

//region MethodRequestContext

// MethodRequestContext provides an abstraction for all input contexts
// in method call inputs.
type MethodRequestContext interface {
	MethodCallContext

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

var _ MethodCallContext = MethodRequestContextBase{}
var _ MethodRequestContext = MethodRequestContextBase{}

// AZContext is required
// for conformance with Context.
func (MethodRequestContextBase) AZContext() {}

// AZMethodCallContext is required
// for conformance with MethodCallContext.
func (MethodRequestContextBase) AZMethodCallContext() {}

// AZMethodRequestContext is required
// for conformance with MethodRequestContext.
func (MethodRequestContextBase) AZMethodRequestContext() {}

//endregion

//region MethodResponseContext

// MethodResponseContext provides an abstraction for all output contexts
// in method call outputs.
//
//TODO: listing of affected states with their respective revision ID.
//TODO: directive: done/end, retry (on failure; optionally with timing and
// retry count parameters or exponentially back-off parameters), redirect
type MethodResponseContext interface {
	MethodCallContext

	AZMethodResponseContext()

	// // Succeed returns true when the method achieved its objective.
	// Succeed() bool

	// Returns the error, if any.
	Err() MethodCallError

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
	err     MethodCallError
	mutated bool
}

var _ MethodResponseContext = MethodResponseContextBase{}

// NewMethodResponseContext creates a new instance
// of MethodResponseContext.
func NewMethodResponseContext(
	err MethodCallError, mutated bool,
) MethodResponseContextBase {
	return MethodResponseContextBase{err: err, mutated: mutated}
}

// AZContext is required for conformance with Context.
func (MethodResponseContextBase) AZContext() {}

// AZMethodCallContext is required
// for conformance with MethodCallContext.
func (MethodResponseContextBase) AZMethodCallContext() {}

// AZMethodResponseContext is required
// for conformance with MethodResponseContext.
func (MethodResponseContextBase) AZMethodResponseContext() {}

// Err is required for conformance with MethodResponseContext.
func (ctx MethodResponseContextBase) Err() MethodCallError { return ctx.err }

// Mutated is required for conformance with MethodResponseContext.
func (ctx MethodResponseContextBase) Mutated() bool { return ctx.mutated }

//endregion