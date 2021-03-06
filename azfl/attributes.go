package azcore

// Attributes abstracts attributes.
type Attributes interface {
	Equatable

	AZAttributes()
}
