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
	AZAdjunctEntityAttributes() AdjunctEntityAttributes
}

// AdjunctEntityAttributesBase is a base
// for AdjunctEntityAttributes implementations.
type AdjunctEntityAttributesBase struct {
}

var _ AdjunctEntityAttributes = AdjunctEntityAttributesBase{}

// AZAdjunctEntityAttributes is required for conformance
// with AdjunctEntityAttributes.
func (attrs AdjunctEntityAttributesBase) AZAdjunctEntityAttributes() AdjunctEntityAttributes {
	return attrs
}
