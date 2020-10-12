package azcore

// RefKey defines the contract for all its concrete implementations.
//
// A RefKey is analogous to path in URL.
//
//     https://example.com/stores/12345/items/456789
//
// The part `/stores/12345/items/456789` is item's ref-key while 456789 is the
// id.
//
// An ID is always scoped, while a ref-key is always unique system-wide.
//
//     https://example.com/stores/StoreA/items/456789
//     https://example.com/stores/StoreB/items/456789
//
// In above example, two items from two different stores have the same ID but
// their ref-keys will not be the same. This system allows each store to scale
// independently without significantly affection each others.
type RefKey interface {
	Equatable

	AZRefKey()

	// ID returns only the ID-part of this ref-key.
	ID() EID

	// RefKeyString returns a string representation of the instance.
	RefKeyString() string
}

// RefKeyFromString is a function which creates an instance of RefKey
// from an input string.
type RefKeyFromString func(refKeyString string) (RefKey, Error)
