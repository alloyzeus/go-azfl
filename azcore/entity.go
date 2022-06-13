package azcore

import "github.com/alloyzeus/go-azfl/azid"

// Entity defines the contract for all its concrete implementations.
//
// For now, this is unused.
type Entity interface {
	AZEntity()
}

// An EntityID is an identifier of an entity.
type EntityID[IDNumT EntityIDNum] interface {
	azid.ID[IDNumT]

	AZEntityID()
}

// An EntityAttributes instance contains the actual attributes of an entity.
// It's on itself is a value object and does not have any identity.
//
// An EntityAttributes instance doesn't hold the ID of its entity instance.
// For the structure that holds both the ID and its attributes, see
// KeyedEntityAttributes, which pratically contains a pair of attributes --
// the ID of the entity and the of attributes of the entity.
type EntityAttributes interface {
	Attributes

	AZEntityAttributes()
}

// KeyedEntityAttributes is a self-identifying data structure that contains
// both the ID of the entity and its representing attributes.
//
//TODO: an envelope with EntityInstanceInfo?
type KeyedEntityAttributes[
	EntityIDNumT EntityIDNum,
	EntityIDT EntityID[EntityIDNumT],
	EntityAttributesT EntityAttributes,
] struct {
	ID   EntityIDT
	Data EntityAttributesT
}

type EntityIDNumMethods interface {
	AZEntityIDNum()
}

// EntityIDNum is the unique or local part of an entity identifier.
//
//TODO: this is a value-object.
type EntityIDNum interface {
	azid.IDNum

	EntityIDNumMethods
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

// EntityInstanceInfoBase is a base implementation of EntityInstanceInfo with
// all attributes are public.
type EntityInstanceInfoBase[
	RevisionNumberT EntityRevisionNumber,
	DeletionInfoT EntityDeletionInfo,
] struct {
	RevisionNumber_ RevisionNumberT
	Deletion_       *DeletionInfoT
}

var _ EntityInstanceInfo[
	int32, EntityDeletionInfoBase,
] = EntityInstanceInfoBase[int32, EntityDeletionInfoBase]{}

func (instanceInfo EntityInstanceInfoBase[
	RevisionNumberT, DeletionInfoT,
]) RevisionNumber() RevisionNumberT {
	return instanceInfo.RevisionNumber_
}

func (instanceInfo EntityInstanceInfoBase[
	RevisionNumberT, DeletionInfoT,
]) Deletion() *DeletionInfoT {
	return instanceInfo.Deletion_
}

func (instanceInfo EntityInstanceInfoBase[
	RevisionNumberT, DeletionInfoT,
]) IsDeleted() bool {
	if delInfo := instanceInfo.Deletion_; delInfo != nil {
		return (*delInfo).Deleted()
	}
	return false
}

type EntityRevisionNumber interface {
	int16 | int32 | int64
}
