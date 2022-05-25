package azcore

import "github.com/alloyzeus/go-azfl/azfl/azid"

// Entity defines the contract for all its concrete implementations.
type Entity interface {
	AZEntity()
}

type EntityIDNumMethods interface {
	AZEntityIDNum()
}

// EntityIDNum defines the contract for all its concrete implementations.
//
//TODO: this is a value-object.
type EntityIDNum interface {
	azid.IDNum

	EntityIDNumMethods
}

// EntityRefKey defines the contract for all its concrete implementations.
type EntityRefKey[IDNumT EntityIDNum] interface {
	azid.RefKey[IDNumT]

	AZEntityRefKey()
}

// EntityAttributes abstracts entity attributes.
type EntityAttributes interface {
	Attributes

	AZEntityAttributes()
}
