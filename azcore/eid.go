package azcore

// EID abstracts entity and entity-like object IDs.
type EID interface {
	Equatable

	AZEID()

	// AZEIDString returns a string representation of the instance.
	//
	//TODO: define what this is for.
	AZEIDString() string
}
