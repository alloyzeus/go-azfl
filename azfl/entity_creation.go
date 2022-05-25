package azcore

// EntityCreationInfo holds information about the creation of an entity.
type EntityCreationInfo[
	SessionIDNumT SessionIDNum, TerminalIDNumT TerminalIDNum, UserIDNumT UserIDNum,
] interface {
	OperationInfo[SessionIDNumT, TerminalIDNumT, UserIDNumT]
}

// EntityCreationInfoBase is the base for all entity creation info.
type EntityCreationInfoBase struct{}

// EntityCreationEvent is the abstraction for all entity creation events.
type EntityCreationEvent[
	SessionIDNumT SessionIDNum, TerminalIDNumT TerminalIDNum, UserIDNumT UserIDNum,
] interface {
	AZEntityCreationEvent()

	CreationInfo() EntityCreationInfo[SessionIDNumT, TerminalIDNumT, UserIDNumT]
}

// EntityCreationEventBase is the base implementation of EntityCreationEvent.
type EntityCreationEventBase struct {
}

//TODO: use tests to assert partial interface implementations
//var _ EntityCreationEvent = EntityCreationEventBase{}

// AZEntityCreationEvent is required for conformance with EntityCreationEvent.
func (EntityCreationEventBase) AZEntityCreationEvent() {}

// EntityCreationRequestContext is the abstraction for all entity creation
// call input contexts.
type EntityCreationRequestContext[
	SessionIDNumT SessionIDNum, TerminalIDNumT TerminalIDNum, UserIDNumT UserIDNum,
] interface {
	EntityMutatingRequestContext[SessionIDNumT, TerminalIDNumT, UserIDNumT]

	AZEntityCreationRequestContext()
}

// EntityCreationResponseContext is the abstraction for all entity creation
// call output contexts.
type EntityCreationResponseContext interface {
	EntityMutatingResponseContext

	AZEntityCreationResponseContext()
}

// EntityCreationResponseContextBase is the base implementation
// for EntityCreationResponseContext.
type EntityCreationResponseContextBase struct {
	ServiceMethodResponseContextBase
}

var _ EntityCreationResponseContext = EntityCreationResponseContextBase{}

// AZEntityCreationResponseContext is required for conformance
// with EntityCreationResponseContext.
func (EntityCreationResponseContextBase) AZEntityCreationResponseContext() {}
