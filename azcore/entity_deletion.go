package azcore

// EntityDeletionInfo provides information about the deletion of an entity
// instance.
type EntityDeletionInfo interface {
	// Deleted returns true when the instance is positively deleted.
	Deleted() bool
}

// EntityDeletionInfoBase is a base implementation of EntityDeletionInfo with
// all attributes are public.
type EntityDeletionInfoBase struct {
	Deleted_ bool
}

// Deleted conforms EntityDeletionInfo interface.
func (deletionInfo EntityDeletionInfoBase) Deleted() bool {
	return deletionInfo.Deleted_
}
