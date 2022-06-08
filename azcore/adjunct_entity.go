package azcore

import "github.com/alloyzeus/go-azfl/azfl/azid"

type AdjunctEntityIDNumMethods interface {
	AZAdjunctEntityIDNum()
}

// AdjunctEntityIDNum abstracts adjunct entity IDs.
type AdjunctEntityIDNum interface {
	azid.IDNum

	AdjunctEntityIDNumMethods
}

// AdjunctEntityRefKey abstracts adjunct entity ref keys.
type AdjunctEntityRefKey[IDNumT AdjunctEntityIDNum] interface {
	azid.RefKey[IDNumT]

	AZAdjunctEntityRefKey()
}

// AdjunctEntityAttributes abstracts adjunct entity attributes.
type AdjunctEntityAttributes interface {
	Attributes

	AZAdjunctEntityAttributes()
}
