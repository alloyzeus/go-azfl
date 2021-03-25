package azcore

import "github.com/alloyzeus/go-azfl/azfl/azid"

// Entity defines the contract for all its concrete implementations.
type Entity interface {
	AZEntity()
}

// EntityIDNum defines the contract for all its concrete implementations.
//
//TODO: this is a value-object.
type EntityIDNum interface {
	azid.IDNum

	AZEntityIDNum()
}

// EntityRefKey defines the contract for all its concrete implementations.
type EntityRefKey interface {
	azid.RefKey

	AZEntityRefKey()
}

// EntityAttributes abstracts entity attributes.
type EntityAttributes interface {
	Attributes

	AZEntityAttributes()
}
