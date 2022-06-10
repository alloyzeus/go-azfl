package azid

import (
	"github.com/alloyzeus/go-azfl/azob"
)

// ID is abstract for entity IDs.
//
// In relation to ID-num, an ID is analogous to path in URL.
//
//     https://example.com/stores/12345/items/456789
//
// The part `/stores/12345/items/456789` is item's ID, and `456789`` is
// the ID-num.
//
// An ID-num is always scoped, while an ID could be used to distinctively
// identify an instance of an entity system-wide.
//
//     https://example.com/stores/StoreABC/items/456789
//     https://example.com/stores/StoreXYZ/items/456789
//
// In above example, two items from two different stores have the same ID-num
// but their respective IDs will not be the same. This system allows each
// store to scale independently without significantly affecting each others.
type ID[IDNumT IDNum] interface {
	azob.Equatable

	BinFieldMarshalable
	BinMarshalable
	TextMarshalable

	// // Returns an array of the hosts' IDs.
	// Hosts() []ID

	// AZIDNum returns only the ID-num part of this ID.
	AZIDNum() IDNumT
}

// IDFromString is a function which creates an instance of ID
// from an input string.
type IDFromString[IDNumT IDNum] func(idString string) (ID[IDNumT], Error)
