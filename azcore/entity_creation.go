package azcore

// EntityCreationInfo holds information about the creation of an entity.
type EntityCreationInfo[
	SessionIDNumT SessionIDNum, SessionRefKeyT SessionRefKey[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalRefKeyT TerminalRefKey[TerminalIDNumT],
	UserIDNumT UserIDNum, UserRefKeyT UserRefKey[UserIDNumT],
	SessionSubjectT SessionSubject[
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT],
	SessionT Session[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
		SessionSubjectT],
] interface {
	OperationInfo[
		SessionIDNumT, SessionRefKeyT, TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
		SessionSubjectT,
		SessionT]
}

// EntityCreationInfoBase is the base for all entity creation info.
type EntityCreationInfoBase struct{}

// EntityCreationEvent is the abstraction for all entity creation events.
type EntityCreationEvent[
	SessionIDNumT SessionIDNum, SessionRefKeyT SessionRefKey[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalRefKeyT TerminalRefKey[TerminalIDNumT],
	UserIDNumT UserIDNum, UserRefKeyT UserRefKey[UserIDNumT],
	SessionSubjectT SessionSubject[
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT],
	SessionT Session[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
		SessionSubjectT],
	EntityCreationInfoT EntityCreationInfo[
		SessionIDNumT, SessionRefKeyT, TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT, SessionSubjectT, SessionT],
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
	SessionIDNumT SessionIDNum, SessionRefKeyT SessionRefKey[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalRefKeyT TerminalRefKey[TerminalIDNumT],
	UserIDNumT UserIDNum, UserRefKeyT UserRefKey[UserIDNumT],
	SessionSubjectT SessionSubject[
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT],
	SessionT Session[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
		SessionSubjectT],
	ServiceMethodCallInputContextT ServiceMethodCallInputContext[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
		SessionSubjectT,
		SessionT],
] interface {
	//TODO: creation is not mutation
	EntityMutatingMethodCallContext[
		SessionIDNumT, SessionRefKeyT, TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT, SessionSubjectT, SessionT,
		ServiceMethodCallInputContextT]

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
