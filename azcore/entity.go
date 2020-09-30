package azcore

// Entity defines the contract for all its concrete implementations.
type Entity interface {
	AZEntity() Entity
}

// EntityID defines the contract for all its concrete implementations.
//
//TODO: this is a value-object.
type EntityID interface {
	AZEntityID() EntityID
}

// EntityRefKey defines the contract for all its concrete implementations.
type EntityRefKey interface {
	RefKey

	AZEntityRefKey() EntityRefKey
}

// EntityEvent defines the contract for all event types of the entity.
type EntityEvent interface {
	Event

	AZEntityEvent() EntityEvent
}

// EntityEventBase provides a basic implementation for all Entity events.
type EntityEventBase struct {
	EventBase
}

var (
	_ EntityEvent = EntityEventBase{}
	_ Event       = EntityEventBase{}
)

// AZEntityEvent is required by EntityEvent.
func (evt EntityEventBase) AZEntityEvent() EntityEvent { return evt }

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

/**/ /**/

// EntityService provides an abstraction for all entity services. This
// abstraction is used by both client and server.
type EntityService interface {
	AZEntityService() EntityService
}

// EntityServiceBase provides a basic implementation for EntityService.
type EntityServiceBase struct {
}

var _ EntityService = &EntityServiceBase{}

// AZEntityService is required for conformance with EntityService.
func (svc *EntityServiceBase) AZEntityService() EntityService { return svc }

/**/ /**/

// EntityServiceClient provides an abstraction for all entity service clients.
type EntityServiceClient interface {
	EntityService

	AZEntityServiceClient() EntityServiceClient
}

/**/ /**/

// EntityServiceServer provides an abstraction for all entity service servers.
type EntityServiceServer interface {
	EntityService

	AZEntityServiceServer() EntityServiceServer
}

/**/ /**/

// EntityCreationEvent is the abstraction for all entity creation events.
type EntityCreationEvent interface {
	AZEntityCreationEvent() EntityCreationEvent
}

// EntityCreationEventBase is the base implementation of EntityCreationEvent.
type EntityCreationEventBase struct{}

var _ EntityCreationEvent = EntityCreationEventBase{}

// AZEntityCreationEvent is required for conformance with EntityCreationEventBase.
func (evt EntityCreationEventBase) AZEntityCreationEvent() EntityCreationEvent {
	return evt
}

// EntityCreationInputContext is the abstraction for all entity creation
// call input contexts.
type EntityCreationInputContext interface {
	EntityMutatingInputContext

	AZEntityCreationInputContext() EntityCreationInputContext
}

// EntityCreationInputContextBase is the base implementation
// for EntityCreationInputContext.
type EntityCreationInputContextBase struct{}

var _ EntityCreationInputContext = EntityCreationInputContextBase{}

// AZEntityCreationInputContext is required for conformance
// with EntityCreationInputContext.
func (evt EntityCreationInputContextBase) AZEntityCreationInputContext() EntityCreationInputContext {
	return evt
}

// EntityCreationOutputContext is the abstraction for all entity creation
// call output contexts.
type EntityCreationOutputContext interface {
	EntityMutatingOutputContext

	AZEntityCreationOutputContext() EntityCreationOutputContext
}

// EntityCreationOutputContextBase is the base implementation
// for EntityCreationOutputContext.
type EntityCreationOutputContextBase struct {
	MethodCallOutputContextBase
}

var _ EntityCreationOutputContext = EntityCreationOutputContextBase{}

// AZEntityCreationOutputContext is required for conformance
// with EntityCreationOutputContext.
func (evt EntityCreationOutputContextBase) AZEntityCreationOutputContext() EntityCreationOutputContext {
	return evt
}

// EntityCreationInfoBase is the base for all entity creation info.
type EntityCreationInfoBase struct{}
