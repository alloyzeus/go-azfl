package azcore

// EntityEvent defines the contract for all event types of the entity.
type EntityEvent interface {
	Event
}

// EntityEventBase provides a basic implementation for all Entity events.
type EntityEventBase struct {
	EventBase
}

var (
	_ EntityEvent = EntityEventBase{}
	_ Event       = EntityEventBase{}
)
