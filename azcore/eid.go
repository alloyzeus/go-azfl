package azcore

// EID abstracts entity and entity-like object IDs.
type EID interface {
	Equatable

	AZEID()
}
