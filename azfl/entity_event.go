package azcore

// EntityEvent defines the contract for all event types of the entity.
type EntityEvent interface {
	Event

	AZEntityEvent()
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
func (EntityEventBase) AZEntityEvent() {}
