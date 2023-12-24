package azcore

// Event provides a contract for all events in the system.
type Event interface {
}

// EventBase provides a common implementation for all events.
type EventBase struct{}

var _ Event = EventBase{}
