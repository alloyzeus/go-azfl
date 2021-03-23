package azcore

// AdjunctEntityID abstracts adjunct entity IDs.
type AdjunctEntityID interface {
	IDNum

	AZAdjunctEntityID()
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
