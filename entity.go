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

// EntityOperationContext provides an abstraction for all operations which
// apply to entity instances.
type EntityOperationContext interface {
	Context
}

// EntityMutatingContext is a specialization of EntityOperationContext which
// is used for operations which make any change to the entity.
type EntityMutatingContext interface {
	EntityOperationContext
}

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

// EntityServiceClient provides an abstraction for all entity service clients.
type EntityServiceClient interface {
	EntityService

	AZEntityServiceClient() EntityServiceClient
}

// EntityServiceServer provides an abstraction for all entity service servers.
type EntityServiceServer interface {
	EntityService

	AZEntityServiceServer() EntityServiceServer
}
