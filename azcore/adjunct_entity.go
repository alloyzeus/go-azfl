package azcore

import "github.com/alloyzeus/go-azfl/azid"

// An AdjunctEntityID is an identifier of an adjunt-entity.
type AdjunctEntityID[IDNumT AdjunctEntityIDNum] interface {
	azid.ID[IDNumT]

	AZAdjunctEntityID()
}

// AdjunctEntityAttributes abstracts adjunct entity attributes.
type AdjunctEntityAttributes interface {
	Attributes

	AZAdjunctEntityAttributes()
}

type AdjunctEntityIDNumMethods interface {
	AZAdjunctEntityIDNum()
}

// AdjunctEntityIDNum abstracts adjunct entity IDs.
type AdjunctEntityIDNum interface {
	azid.IDNum

	AdjunctEntityIDNumMethods
}
