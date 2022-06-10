package azcore

type EntityDeletionInfo interface {
	Deleted() bool
}

type EntityDeletionInfoBase struct {
	Deleted_ bool
}

func (deletionInfo EntityDeletionInfoBase) Deleted() bool { return deletionInfo.Deleted_ }
