package azcore

//----

// MethodCallError provides an abstraction for all errors returned by a method.
//
//TODO: sub-classes: input (acces, parameters, context), internal
type MethodCallError interface {
	Error

	AZMethodCallError()
}

//----

// MethodCallContext is an abstraction for input and output contexts used
// when calling a method.
type MethodCallContext interface {
	Context

	AZMethodCallContext()
}

//----

// MethodCallID represents the identifier of a method call. This identifier
// doubles as idempotency token.
type MethodCallID interface {
	Equatable
	AZMethodCallID()
}

// MethodCallInputContext provides an abstraction for all input contexts
// in method call inputs.
type MethodCallInputContext interface {
	MethodCallContext

	AZMethodCallInputContext()
}

//----

// MethodCallInputContextBase is a partial implementation
// of MethodCallInputContext.
type MethodCallInputContextBase struct{}

var _ MethodCallContext = MethodCallInputContextBase{}
var _ MethodCallInputContext = MethodCallInputContextBase{}

// AZContext is required
// for conformance with Context.
func (MethodCallInputContextBase) AZContext() {}

// AZMethodCallContext is required
// for conformance with MethodCallContext.
func (MethodCallInputContextBase) AZMethodCallContext() {}

// AZMethodCallInputContext is required
// for conformance with MethodCallInputContext.
func (MethodCallInputContextBase) AZMethodCallInputContext() {}

//----

// MethodCallOutputContext provides an abstraction for all output contexts
// in method call outputs.
type MethodCallOutputContext interface {
	MethodCallContext

	AZMethodCallOutputContext()

	// Returns the error, if any.
	Err() MethodCallError

	// // Failed returns true when the method unable to fulfill its
	// // main objective.
	// Failed() bool

	// Mutated returns true if the method made any changes, even when the
	// method failed to fulfill its objective. It should not
	// return true if the change has been completely rolled-back before the
	// method returned this context.
	Mutated() bool
}

//----

// MethodCallOutputContextBase is a base
// for MethodCallOutputContext implementations.
type MethodCallOutputContextBase struct {
	err     MethodCallError
	mutated bool
}

var _ MethodCallOutputContext = MethodCallOutputContextBase{}

// NewMethodCallOutputContext creates a new instance
// of MethodCallOutputContext.
func NewMethodCallOutputContext(
	err MethodCallError, mutated bool,
) MethodCallOutputContextBase {
	return MethodCallOutputContextBase{err: err, mutated: mutated}
}

// AZContext is required for conformance with Context.
func (MethodCallOutputContextBase) AZContext() {}

// AZMethodCallContext is required
// for conformance with MethodCallContext.
func (MethodCallOutputContextBase) AZMethodCallContext() {}

// AZMethodCallOutputContext is required
// for conformance with MethodCallOutputContext.
func (MethodCallOutputContextBase) AZMethodCallOutputContext() {}

// Err is required for conformance with MethodCallOutputContext.
func (outputCtx MethodCallOutputContextBase) Err() MethodCallError { return outputCtx.err }

// Mutated is required for conformance with MethoCallOutputContext.
func (outputCtx MethodCallOutputContextBase) Mutated() bool { return outputCtx.mutated }
