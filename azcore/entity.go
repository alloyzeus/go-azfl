package azcore

import "github.com/alloyzeus/go-azfl/azid"

// Entity defines the contract for all its concrete implementations.
type Entity interface {
	AZEntity()
}

// EntityEnvelope is a self-identifying data structure that contains the
// id of the entity and its representing data.
type EntityEnvelope[
	EntityIDNumT EntityIDNum,
	EntityRefKeyT EntityRefKey[EntityIDNumT],
	EntityDataT EntityData,
] struct {
}

type EntityData interface{}

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
