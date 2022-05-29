package azcore

//region Method

// EntityMethodMessage abstracts the messages, i.e., requests and responses.
type EntityMethodMessage interface {
	ServiceMethodMessage

	EntityMethodContext() EntityMethodContext
}

// EntityMethodRequest abstracts all entity method requests messages.
type EntityMethodRequest[
	SessionIDNumT SessionIDNum, TerminalIDNumT TerminalIDNum, UserIDNumT UserIDNum,
	EntityMethodRequestContextT EntityMethodRequestContext[SessionIDNumT, TerminalIDNumT, UserIDNumT],
] interface {
	ServiceMethodRequest[SessionIDNumT, TerminalIDNumT, UserIDNumT]
	EntityMethodMessage

	EntityMethodRequestContext() EntityMethodRequestContextT
}

// EntityMethodResponse abstracts all entity method response messages.
type EntityMethodResponse interface {
	ServiceMethodResponse
	EntityMethodMessage

	EntityMethodResponseContext() EntityMethodResponseContext
}

//endregion

//region Context

// EntityMethodContext provides an abstraction for all operations which
// apply to entity instances.
type EntityMethodContext interface {
	ServiceMethodContext
}

// EntityMethodRequestContext is an abstraction for all method call
// input contexts.
type EntityMethodRequestContext[
	SessionIDNumT SessionIDNum, TerminalIDNumT TerminalIDNum, UserIDNumT UserIDNum,
] interface {
	EntityMethodContext
	ServiceMethodRequestContext[
		SessionIDNumT, TerminalIDNumT, UserIDNumT, Session[
			SessionIDNumT, SessionRefKey[SessionIDNumT],
			TerminalIDNumT, TerminalRefKey[TerminalIDNumT],
			UserIDNumT, UserRefKey[UserIDNumT],
			Subject[
				TerminalIDNumT, TerminalRefKey[TerminalIDNumT],
				UserIDNumT, UserRefKey[UserIDNumT]],
		],
	]
}

// EntityMethodResponseContext is an abstraction for all method call
// output contexts.
type EntityMethodResponseContext interface {
	EntityMethodContext
	ServiceMethodResponseContext
}

//endregion

//region MutatingContext

// EntityMutatingContext is a specialization of EntityOperationContext which
// is used for operations which make any change to the entity.
type EntityMutatingContext interface {
	EntityMethodContext
	ServiceMutatingMethodContext
}

// EntityMutatingRequestContext provides an abstraction for input contexts
// for mutating method calls.
type EntityMutatingRequestContext[
	SessionIDNumT SessionIDNum, TerminalIDNumT TerminalIDNum, UserIDNumT UserIDNum,
] interface {
	EntityMutatingContext
	EntityMethodRequestContext[SessionIDNumT, TerminalIDNumT, UserIDNumT]
	ServiceMutatingOpRequestContext[SessionIDNumT, TerminalIDNumT, UserIDNumT]
}

// EntityMutatingResponseContext provides an abstraction for output contexts
// for mutating method calls.
type EntityMutatingResponseContext interface {
	EntityMutatingContext
	EntityMethodResponseContext
	ServiceMutatingMethodResponseContext
}

// EntityMutatingMessage abstracts entity mutating method requests and responses.
type EntityMutatingMessage interface {
	EntityMethodMessage
	ServiceMutatingMethodMessage

	EntityMutatingContext() EntityMutatingContext
}

// EntityMutatingRequest abstracts entity mutating requests.
type EntityMutatingRequest[
	SessionIDNumT SessionIDNum, TerminalIDNumT TerminalIDNum, UserIDNumT UserIDNum,
	EntityMutatingRequestContextT EntityMutatingRequestContext[SessionIDNumT, TerminalIDNumT, UserIDNumT],
] interface {
	EntityMutatingMessage
	ServiceMutatingMethodRequest[SessionIDNumT, TerminalIDNumT, UserIDNumT]

	EntityMutatingRequestContext() EntityMutatingRequestContextT
}

// EntityMutatingResponse abstracts entity mutating responses.
type EntityMutatingResponse interface {
	EntityMutatingMessage
	ServiceMutatingMethodResponse

	EntityMutatingResponseContext() EntityMutatingResponseContext
}

//endregion

//region Service

// EntityService provides an abstraction for all entity services. This
// abstraction is used by both client and server.
type EntityService interface {
	AZEntityService()
}

//endregion

//region ServiceBase

// EntityServiceBase provides a basic implementation for EntityService.
type EntityServiceBase struct{}

var _ EntityService = &EntityServiceBase{}

// AZEntityService is required for conformance with EntityService.
func (*EntityServiceBase) AZEntityService() {}

//endregion

//region ServiceClient

// EntityServiceClient provides an abstraction for all entity service clients.
type EntityServiceClient interface {
	EntityService

	AZEntityServiceClient()
}

//endregion

//region ServiceServer

// EntityServiceServer provides an abstraction for all entity service servers.
type EntityServiceServer interface {
	EntityService

	AZEntityServiceServer()
}

//endregion
