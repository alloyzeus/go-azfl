package azcore

// AdjunctEntityID abstracts adjunct entity IDs.
type AdjunctEntityID interface {
	AZAdjunctEntityID() AdjunctEntityID
}

// AdjunctEntityRefKey abstracts adjunct entity ref keys.
type AdjunctEntityRefKey interface {
	AZAdjunctEntityRefKey() AdjunctEntityRefKey
}

// AdjunctEntityAttributes abstracts adjunct entity attributes.
type AdjunctEntityAttributes interface {
	Equatable

	AZAdjunctEntityAttributes() AdjunctEntityAttributes
}
