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

//----

// EntityAttributes abstracts entity attributes.
type EntityAttributes interface {
	Equatable

	AZEntityAttributes() EntityAttributes
}

//----

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

// EntityCreationInfoBase is the base for all entity creation info.
type EntityCreationInfoBase struct{}
