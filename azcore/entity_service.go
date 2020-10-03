package azcore

//region CallContext

// EntityMethodCallContext provides an abstraction for all operations which
// apply to entity instances.
type EntityMethodCallContext interface {
	MethodCallContext
}

// EntityMethodCallInputContext is an abstraction for all method call
// input contexts.
type EntityMethodCallInputContext interface {
	EntityMethodCallContext
	MethodCallInputContext
}

// EntityMethodCallOutputContext is an abstraction for all method call
// output contexts.
type EntityMethodCallOutputContext interface {
	EntityMethodCallContext
	MethodCallOutputContext
}

//endregion

//region MutatingCallContext

// EntityMutatingCallContext is a specialization of EntityOperationContext which
// is used for operations which make any change to the entity.
type EntityMutatingCallContext interface {
	EntityMethodCallContext
}

// EntityMutatingInputContext provides an abstraction for input contexts
// for mutating method calls.
type EntityMutatingInputContext interface {
	EntityMutatingCallContext
	EntityMethodCallInputContext
}

// EntityMutatingOutputContext provides an abstraction for output contexts
// for mutating method calls.
type EntityMutatingOutputContext interface {
	EntityMutatingCallContext
	EntityMethodCallOutputContext
}

//endregion

//region Service

// EntityService provides an abstraction for all entity services. This
// abstraction is used by both client and server.
type EntityService interface {
	AZEntityService()
}

//endregion

//region ServiceBase

// EntityServiceBase provides a basic implementation for EntityService.
type EntityServiceBase struct{}

var _ EntityService = &EntityServiceBase{}

// AZEntityService is required for conformance with EntityService.
func (*EntityServiceBase) AZEntityService() {}

//endregion

//region ServiceClient

// EntityServiceClient provides an abstraction for all entity service clients.
type EntityServiceClient interface {
	EntityService

	AZEntityServiceClient()
}

//endregion

//region ServiceServer

// EntityServiceServer provides an abstraction for all entity service servers.
type EntityServiceServer interface {
	EntityService

	AZEntityServiceServer()
}

//endregion
