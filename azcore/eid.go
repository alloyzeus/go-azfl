package azcore

import "github.com/alloyzeus/go-azcore/azcore/azwire"

// EID abstracts entity and entity-like object IDs.
//
//TODO: define that an EID must be of a primitive type.
type EID interface {
	Equatable

	AZEID()

	// An EID must be azwire-marshalable as a field. It never need to be
	// marshalable as a top-level object.
	azwire.FieldMarshalable
}
