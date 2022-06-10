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
	ID   EntityIDT
	Data EntityDataT
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

// EntityInstanceInfo holds information about an instance of entity, i.e.,
// metadata of an instance of entity. It doesn't contain the attributes of
// the instance itself.
type EntityInstanceInfo[
	RevisionNumberT EntityRevisionNumber,
	DeletionInfoT EntityDeletionInfo,
] interface {
	RevisionNumber() RevisionNumberT

	// Deletion returns a detailed information about the deletion if
	// the instance has been deleted.
	Deletion() *DeletionInfoT
	// IsDeleted returns true if the instance has been deleted.
	IsDeleted() bool
}

type EntityInstanceInfoBase[
	RevisionNumberT EntityRevisionNumber,
	DeletionInfoT EntityDeletionInfo,
] struct {
	RevisionNumber_ RevisionNumberT
	Deletion_       *DeletionInfoT
}

var _ EntityInstanceInfo[int32, EntityDeletionInfoBase] = EntityInstanceInfoBase[int32, EntityDeletionInfoBase]{}

func (instanceInfo EntityInstanceInfoBase[RevisionNumberT, DeletionInfoT]) RevisionNumber() RevisionNumberT {
	return instanceInfo.RevisionNumber_
}

func (entityInstanceInfo EntityInstanceInfoBase[RevisionNumberT, DeletionInfoT]) Deletion() *DeletionInfoT {
	return entityInstanceInfo.Deletion_
}

func (entityInstanceInfo EntityInstanceInfoBase[RevisionNumberT, DeletionInfoT]) IsDeleted() bool {
	return entityInstanceInfo.Deletion_ != nil && (*entityInstanceInfo.Deletion_).Deleted()
}

type EntityRevisionNumber interface {
	int16 | int32 | int64
}
