package azcore

// EntityCreationInfo holds information about the creation of an entity.
type EntityCreationInfo interface {
	ActionInfo
}

// EntityCreationInfoBase is the base for all entity creation info.
type EntityCreationInfoBase struct{}

// EntityCreationEvent is the abstraction for all entity creation events.
type EntityCreationEvent interface {
	AZEntityCreationEvent()
}

// EntityCreationEventBase is the base implementation of EntityCreationEvent.
type EntityCreationEventBase struct{}

var _ EntityCreationEvent = EntityCreationEventBase{}

// AZEntityCreationEvent is required for conformance with EntityCreationEventBase.
func (EntityCreationEventBase) AZEntityCreationEvent() {}

// EntityCreationInputContext is the abstraction for all entity creation
// call input contexts.
type EntityCreationInputContext interface {
	EntityMutatingInputContext

	AZEntityCreationInputContext()
}

// EntityCreationInputContextBase is the base implementation
// for EntityCreationInputContext.
type EntityCreationInputContextBase struct {
	MethodCallInputContextBase
}

var _ EntityCreationInputContext = EntityCreationInputContextBase{}

// AZEntityCreationInputContext is required for conformance
// with EntityCreationInputContext.
func (EntityCreationInputContextBase) AZEntityCreationInputContext() {}

// EntityCreationOutputContext is the abstraction for all entity creation
// call output contexts.
type EntityCreationOutputContext interface {
	EntityMutatingOutputContext

	AZEntityCreationOutputContext()
}

// EntityCreationOutputContextBase is the base implementation
// for EntityCreationOutputContext.
type EntityCreationOutputContextBase struct {
	MethodCallOutputContextBase
}

var _ EntityCreationOutputContext = EntityCreationOutputContextBase{}

// AZEntityCreationOutputContext is required for conformance
// with EntityCreationOutputContext.
func (EntityCreationOutputContextBase) AZEntityCreationOutputContext() {}
