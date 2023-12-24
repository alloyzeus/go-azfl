package azcore

type EntityCreationOutput[
	IDNumT EntityIDNum,
	IDT EntityID[IDNumT],
	RevisionNumberT EntityRevisionNumber,
	DeletionInfoT EntityDeletionInfo,
	InstanceInfoT EntityInstanceInfo[
		RevisionNumberT, DeletionInfoT],
] struct {
	InstanceID   IDT
	InitialState InstanceInfoT
}

// EntityCreationInfo holds information about the creation of an entity.
type EntityCreationInfo[
	SessionIDNumT SessionIDNum, SessionIDT SessionID[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalIDT TerminalID[TerminalIDNumT],
	UserIDNumT UserIDNum, UserIDT UserID[UserIDNumT],
	SessionSubjectT SessionSubject[
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT],
	SessionT Session[
		SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT, SessionT],
] interface {
	OperationInfo[
		SessionIDNumT, SessionIDT, TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT,
		SessionT]
}

// EntityCreationEvent is the abstraction for all entity creation events.
type EntityCreationEvent[
	SessionIDNumT SessionIDNum, SessionIDT SessionID[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalIDT TerminalID[TerminalIDNumT],
	UserIDNumT UserIDNum, UserIDT UserID[UserIDNumT],
	SessionSubjectT SessionSubject[
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT],
	SessionT Session[
		SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT, SessionT],
	EntityCreationInfoT EntityCreationInfo[
		SessionIDNumT, SessionIDT, TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT, SessionSubjectT, SessionT],
] interface {
	CreationInfo() EntityCreationInfoT
}

// EntityCreationCallContext is the abstraction for all entity creation
// call input contexts.
type EntityCreationCallContext[
	SessionIDNumT SessionIDNum, SessionIDT SessionID[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalIDT TerminalID[TerminalIDNumT],
	UserIDNumT UserIDNum, UserIDT UserID[UserIDNumT],
	SessionSubjectT SessionSubject[
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT],
	SessionT Session[
		SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT, SessionT],
	ServiceOpIdempotencyKeyT ServiceOpIdempotencyKey,
] interface {
	ServiceOpCallContext[
		SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT,
		SessionT, ServiceOpIdempotencyKeyT]
}

// EntityCreationResultContext is the abstraction for all entity creation
// call output contexts.
type EntityCreationResultContext interface {
	ServiceOpResultContext
}
