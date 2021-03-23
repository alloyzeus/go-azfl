package azcore

import (
	"github.com/alloyzeus/go-azfl/azfl/azer"
)

// IDNum abstracts entity and entity-like object IDs.
//
//TODO: define that an IDNum must be of a primitive type.
type IDNum interface {
	Equatable

	AZIDNum()

	// An IDNum must be azer-bin-marshalable as a field. It never need to be
	// marshalable as a top-level object.
	azer.BinFieldMarshalable
}
