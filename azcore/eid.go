package azcore

// EID abstracts entity and entity-like object IDs.
//
//TODO: define that an EID must be of a primitive type.
type EID interface {
	Equatable

	AZEID()

	AZWireObject
}
