package azcore

import "github.com/alloyzeus/go-azfl/azid"

type AdjunctEntityIDNumMethods interface {
	AZAdjunctEntityIDNum()
}

// AdjunctEntityIDNum abstracts adjunct entity IDs.
type AdjunctEntityIDNum interface {
	azid.IDNum

	AdjunctEntityIDNumMethods
}

// AdjunctEntityID abstracts adjunct entity ref keys.
type AdjunctEntityID[IDNumT AdjunctEntityIDNum] interface {
	azid.ID[IDNumT]

	AZAdjunctEntityID()
}

// AdjunctEntityAttributes abstracts adjunct entity attributes.
type AdjunctEntityAttributes interface {
	Attributes

	AZAdjunctEntityAttributes()
}
