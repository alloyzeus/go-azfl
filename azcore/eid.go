package azcore

// EID abstracts entity and entity-like object IDs.
//
//TODO: define that an EID must be of a primitive type.
type EID interface {
	Equatable

	AZEID()

	// AZEIDBinary returns a binary respresentation of the instance.
	// The encoding must be protobuf-compatible.
	AZEIDBinary() []byte

	// AZEIDString returns a string representation of the instance.
	//
	//TODO: define what this is for.
	AZEIDString() string
}
