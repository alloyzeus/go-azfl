package azcore

// RefKey defines the contract for all its concrete implementations.
type RefKey interface {
	AZRefKey() RefKey

	// RefKeyString returns a string representation of the instance.
	RefKeyString() string
}

// RefKeyFromString is a function which creates an instance of RefKey
// from an input string.
type RefKeyFromString func(refKeyString string) (RefKey, Error)
