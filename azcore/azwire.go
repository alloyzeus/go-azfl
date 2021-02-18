package azcore

// AZWire is protowire-conformance encoding.

// AZWireObject provides an interface for objects which could be encoded
// to / decoded from azwire encoding.
type AZWireObject interface {
	// AZWire returns azwire-encoded bytes representing the instance.
	AZWire() []byte

	// UnmarshalAZWire modifies the instance with data from
	// azwire-encoded bytes.
	UnmarshalAZWire(b []byte) error
}
