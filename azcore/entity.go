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
	EntityIDT EntityID[EntityIDNumT],
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

// EntityID defines the contract for all its concrete implementations.
type EntityID[IDNumT EntityIDNum] interface {
	azid.ID[IDNumT]

	AZEntityID()
}

// EntityAttributes abstracts entity attributes.
type EntityAttributes interface {
	Attributes

	AZEntityAttributes()
}
