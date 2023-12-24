package azcore

import "github.com/alloyzeus/go-azfl/v2/azid"

// An AdjunctEntityID is an identifier of an adjunt-entity.
type AdjunctEntityID[IDNumT AdjunctEntityIDNum] interface {
	azid.ID[IDNumT]
}

// AdjunctEntityAttrSet abstracts adjunct entity attributes.
type AdjunctEntityAttrSet interface {
	AttrSet
}

type AdjunctEntityIDNumMethods interface {
}

// AdjunctEntityIDNum abstracts adjunct entity IDs.
type AdjunctEntityIDNum interface {
	azid.IDNum

	AdjunctEntityIDNumMethods
}
