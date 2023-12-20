package azcore

// ServiceMutationOpMessage abstracts mutating method requests
// and responses.
type ServiceMutationOpMessage interface {
	ServiceOpMessage

	MutationOpContext() ServiceMutationOpContext
}

// ServiceMutationOpContext abstracts contexts of mutating
// method requests and responses.
type ServiceMutationOpContext interface {
	ServiceOpContext
}
