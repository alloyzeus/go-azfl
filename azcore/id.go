package azcore

// ID abstracts IDs.
type ID interface {
	Equatable

	AZID()
}
