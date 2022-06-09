package azcore

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

//endregion
