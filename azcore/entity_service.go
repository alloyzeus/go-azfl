package azcore

//region Method

// EntityMethodMessage abstracts the messages, i.e., requests and responses.
type EntityMethodMessage interface {
	MethodMessage

	EntityMethodContext() EntityMethodContext
}

// EntityMethodRequest abstracts all entity method requests messages.
type EntityMethodRequest interface {
	MethodRequest
	EntityMethodMessage

	EntityMethodRequestContext() EntityMethodRequestContext
}

// EntityMethodResponse abstracts all entity method response messages.
type EntityMethodResponse interface {
	MethodResponse
	EntityMethodMessage

	EntityMethodResponseContext() EntityMethodResponseContext
}

//endregion

//region Context

// EntityMethodContext provides an abstraction for all operations which
// apply to entity instances.
type EntityMethodContext interface {
	MethodContext
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

//region MutatingContext

// EntityMutatingContext is a specialization of EntityOperationContext which
// is used for operations which make any change to the entity.
type EntityMutatingContext interface {
	EntityMethodContext
	MutatingContext
}

// EntityMutatingRequestContext provides an abstraction for input contexts
// for mutating method calls.
type EntityMutatingRequestContext interface {
	EntityMutatingContext
	EntityMethodRequestContext
	MutatingRequestContext
}

// EntityMutatingResponseContext provides an abstraction for output contexts
// for mutating method calls.
type EntityMutatingResponseContext interface {
	EntityMutatingContext
	EntityMethodResponseContext
	MutatingResponseContext
}

// EntityMutatingMessage abstracts entity mutating method requests and responses.
type EntityMutatingMessage interface {
	EntityMethodMessage
	MutatingMessage

	EntityMutatingContext() EntityMutatingContext
}

// EntityMutatingRequest abstracts entity mutating requests.
type EntityMutatingRequest interface {
	EntityMutatingMessage
	MutatingRequest

	EntityMutatingRequestContext() EntityMutatingRequestContext
}

// EntityMutatingResponse abstracts entity mutating responses.
type EntityMutatingResponse interface {
	EntityMutatingMessage
	MutatingResponse

	EntityMutatingResponseContext() EntityMutatingResponseContext
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
