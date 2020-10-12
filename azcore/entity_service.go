package azcore

//region Method

// EntityMethodMessage abstracts the messages, i.e., requests and responses.
type EntityMethodMessage interface {
	MethodCallMessage
}

// EntityMethodRequest abstracts all entity method requests messages.
type EntityMethodRequest interface {
	MethodRequest

	EntityMethodRequestContext() EntityMethodRequestContext
}

// EntityMethodResponse abstracts all entity method response messages.
type EntityMethodResponse interface {
	MethodResponse

	EntityMethodResponseContext() EntityMethodResponseContext
}

//endregion

//region Context

// EntityMethodContext provides an abstraction for all operations which
// apply to entity instances.
type EntityMethodContext interface {
	MethodCallContext
}

// EntityMethodRequestContext is an abstraction for all method call
// input contexts.
type EntityMethodRequestContext interface {
	EntityMethodContext
	MethodRequestContext
}

// EntityMethodResponseContext is an abstraction for all method call
// output contexts.
type EntityMethodResponseContext interface {
	EntityMethodContext
	MethodResponseContext
}

//endregion

//region MutatingMethodContext

// EntityMutatingMethodContext is a specialization of EntityOperationContext which
// is used for operations which make any change to the entity.
type EntityMutatingMethodContext interface {
	EntityMethodContext
}

// EntityMutatingRequestContext provides an abstraction for input contexts
// for mutating method calls.
type EntityMutatingRequestContext interface {
	EntityMutatingMethodContext
	EntityMethodRequestContext
}

// EntityMutatingResponseContext provides an abstraction for output contexts
// for mutating method calls.
type EntityMutatingResponseContext interface {
	EntityMutatingMethodContext
	EntityMethodResponseContext
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
