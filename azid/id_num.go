package azid

import (
	"github.com/alloyzeus/go-azfl/azob"
)

type IDNumMethods interface {
	azob.Equatable

	AZIDNum()

	// An IDNum must be azid-bin-marshalable as a field. It never need to be
	// marshalable as a top-level object.
	BinFieldMarshalable
}

// IDNum abstracts entity and entity-like object IDs.
//
//TODO: define that an IDNum must be of a primitive type.
type IDNum interface {
	~int16 | ~int32 | ~int64

	IDNumMethods
}
