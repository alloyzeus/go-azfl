package azcore

// AZWire is protowire-conformance encoding.

// AZWireObject provides an interface for objects which could be encoded
// to / decoded from azwire encoding.
type AZWireObject interface {
	AZWireMarshalable

	// Commented-out as we need to pass around pointers.
	// AZWireUnmarshalable
}

// AZWireMarshalable is an interface for objects which could be marshaled
// with azwire encoding.
type AZWireMarshalable interface {
	// AZWire returns azwire-encoded bytes representing the instance.
	AZWire() []byte
}

// An AZWireUnmarshalable is an object which could load azwire-encoded
// bytes into itself.
type AZWireUnmarshalable interface {
	// UnmarshalAZWire modifies the instance with data from
	// azwire-encoded bytes.
	UnmarshalAZWire(b []byte) error
}
