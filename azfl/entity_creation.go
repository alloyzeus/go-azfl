package azcore

// EntityCreationInfo holds information about the creation of an entity.
type EntityCreationInfo interface {
	OperationInfo
}

// EntityCreationInfoBase is the base for all entity creation info.
type EntityCreationInfoBase struct{}

// EntityCreationEvent is the abstraction for all entity creation events.
type EntityCreationEvent interface {
	AZEntityCreationEvent()

	CreationInfo() EntityCreationInfo
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
type EntityCreationRequestContext interface {
	EntityMutatingRequestContext

	AZEntityCreationRequestContext()
}

// EntityCreationRequestContextBase is the base implementation
// for EntityCreationRequestContext.
type EntityCreationRequestContextBase struct {
	ServiceMethodRequestContextBase
}

var _ EntityCreationRequestContext = EntityCreationRequestContextBase{}

// AZEntityCreationRequestContext is required for conformance
// with EntityCreationRequestContext.
func (EntityCreationRequestContextBase) AZEntityCreationRequestContext() {}

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
