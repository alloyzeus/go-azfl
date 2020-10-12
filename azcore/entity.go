package azcore

// Entity defines the contract for all its concrete implementations.
type Entity interface {
	AZEntity()
}

// EntityID defines the contract for all its concrete implementations.
//
//TODO: this is a value-object.
type EntityID interface {
	EID

	AZEntityID()
}

// EntityRefKey defines the contract for all its concrete implementations.
type EntityRefKey interface {
	RefKey

	AZEntityRefKey()
}

// EntityAttributes abstracts entity attributes.
type EntityAttributes interface {
	Attributes

	AZEntityAttributes()
}
