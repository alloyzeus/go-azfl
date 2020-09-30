package azcore

// MethodCallContext is an abstraction for input and output contexts used
// when calling a method.
type MethodCallContext interface {
	Context
}

// MethodCallInputContext provides an abstraction for all input contexts
// in method call inputs.
type MethodCallInputContext interface {
	MethodCallContext
}

// MethodCallOutputContext provides an abstraction for all output contexts
// in method call outputs.
type MethodCallOutputContext interface {
	MethodCallContext
}
