package azcore

import "github.com/alloyzeus/go-azfl/azfl/azid"

// AdjunctEntityIDNum abstracts adjunct entity IDs.
type AdjunctEntityIDNum interface {
	azid.IDNum

	AZAdjunctEntityIDNum()
}

// AdjunctEntityRefKey abstracts adjunct entity ref keys.
type AdjunctEntityRefKey interface {
	azid.RefKey

	AZAdjunctEntityRefKey()
}

// AdjunctEntityAttributes abstracts adjunct entity attributes.
type AdjunctEntityAttributes interface {
	Attributes

	AZAdjunctEntityAttributes()
}
