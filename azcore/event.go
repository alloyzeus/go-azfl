package azcore

// Event provides a contract for all events in the system.
type Event interface {
	AZEvent()
}

// EventBase provides a common implementation for all events.
type EventBase struct{}

var _ Event = EventBase{}

// AZEvent is required by Event
func (EventBase) AZEvent() {}
