package azcore

// AdjunctEntityID abstracts adjunct entity IDs.
type AdjunctEntityID interface {
	EID

	AZAdjunctEntityID()

	// AZAdjunctEntityIDString returns a string representation of the instance.
	//
	//TODO: define what this is for.
	AZAdjunctEntityIDString() string
}

// AdjunctEntityRefKey abstracts adjunct entity ref keys.
type AdjunctEntityRefKey interface {
	RefKey

	AZAdjunctEntityRefKey()
}

// AdjunctEntityAttributes abstracts adjunct entity attributes.
type AdjunctEntityAttributes interface {
	Attributes

	AZAdjunctEntityAttributes()
}
