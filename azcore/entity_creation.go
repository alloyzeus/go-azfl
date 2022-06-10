package azcore

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
		SessionSubjectT],
] interface {
	OperationInfo[
		SessionIDNumT, SessionIDT, TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT,
		SessionT]
}

// EntityCreationInfoBase is the base for all entity creation info.
type EntityCreationInfoBase struct{}

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
		SessionSubjectT],
	EntityCreationInfoT EntityCreationInfo[
		SessionIDNumT, SessionIDT, TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT, SessionSubjectT, SessionT],
] interface {
	AZEntityCreationEvent()

	CreationInfo() EntityCreationInfoT
}

// EntityCreationEventBase is the base implementation of EntityCreationEvent.
type EntityCreationEventBase struct {
}

//TODO: use tests to assert partial interface implementations
//var _ EntityCreationEvent = EntityCreationEventBase{}

// AZEntityCreationEvent is required for conformance with EntityCreationEvent.
func (EntityCreationEventBase) AZEntityCreationEvent() {}

// EntityCreationInputContext is the abstraction for all entity creation
// call input contexts.
type EntityCreationInputContext[
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
		SessionSubjectT],
	ServiceMethodCallInputContextT ServiceMethodCallInputContext[
		SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT,
		SessionT, ServiceMethodIdempotencyKeyT],
	ServiceMethodIdempotencyKeyT ServiceMethodIdempotencyKey,
] interface {
	//TODO: creation is not mutation
	EntityMutatingMethodCallContext[
		SessionIDNumT, SessionIDT, TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT, SessionSubjectT, SessionT,
		ServiceMethodCallInputContextT, ServiceMethodIdempotencyKeyT]

	AZEntityCreationInputContext()
}

// EntityCreationOutputContext is the abstraction for all entity creation
// call output contexts.
type EntityCreationOutputContext interface {
	EntityMutatingMethodCallOutputContext

	AZEntityCreationOutputContext()
}

// EntityCreationOutputContextBase is the base implementation
// for EntityCreationOutputContext.
type EntityCreationOutputContextBase struct {
	ServiceMethodCallOutputContextBase
}

var _ EntityCreationOutputContext = EntityCreationOutputContextBase{}

// AZEntityCreationOutputContext is required for conformance
// with EntityCreationOutputContext.
func (EntityCreationOutputContextBase) AZEntityCreationOutputContext() {}
