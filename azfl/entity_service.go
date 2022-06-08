package azcore

//region Method

// EntityMethodMessage abstracts the messages, i.e., requests and responses.
type EntityMethodMessage interface {
	ServiceMethodMessage

	EntityMethodContext() EntityMethodContext
}

// EntityMethodRequest abstracts all entity method requests messages.
type EntityMethodRequest[
	SessionIDNumT SessionIDNum, SessionRefKeyT SessionRefKey[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalRefKeyT TerminalRefKey[TerminalIDNumT],
	UserIDNumT UserIDNum, UserRefKeyT UserRefKey[UserIDNumT],
	SessionSubjectT SessionSubject[
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
	],
	SessionT Session[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
		SessionSubjectT,
	],
	ServiceMethodCallInputContextT ServiceMethodCallInputContext[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
		SessionSubjectT,
		SessionT,
	],
	EntityMethodRequestContextT EntityMethodRequestContext[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT, SessionSubjectT, SessionT],
] interface {
	ServiceMethodCallInput[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT, SessionSubjectT, SessionT,
		ServiceMethodCallInputContextT]
	EntityMethodMessage

	EntityMethodRequestContext() EntityMethodRequestContextT
}

// EntityMethodResponse abstracts all entity method response messages.
type EntityMethodResponse interface {
	ServiceMethodCallOutput
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
	SessionIDNumT SessionIDNum, SessionRefKeyT SessionRefKey[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalRefKeyT TerminalRefKey[TerminalIDNumT],
	UserIDNumT UserIDNum, UserRefKeyT UserRefKey[UserIDNumT],
	SessionSubjectT SessionSubject[
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
	],
	SessionT Session[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
		SessionSubjectT,
	],
] interface {
	EntityMethodContext
	ServiceMethodCallInputContext[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
		SessionSubjectT,
		SessionT,
	]
}

// EntityMethodResponseContext is an abstraction for all method call
// output contexts.
type EntityMethodResponseContext interface {
	EntityMethodContext
	ServiceMethodCallOutputContext
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
	SessionIDNumT SessionIDNum, SessionRefKeyT SessionRefKey[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalRefKeyT TerminalRefKey[TerminalIDNumT],
	UserIDNumT UserIDNum, UserRefKeyT UserRefKey[UserIDNumT],
	SessionSubjectT SessionSubject[
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
	],
	SessionT Session[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
		SessionSubjectT,
	],
	ServiceMethodCallInputContextT ServiceMethodCallInputContext[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
		SessionSubjectT,
		SessionT,
	],
] interface {
	EntityMutatingContext
	EntityMethodRequestContext[
		SessionIDNumT, SessionRefKeyT, TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT, SessionSubjectT, SessionT]
	ServiceMutatingMethodCallInputContext[
		SessionIDNumT, SessionRefKeyT, TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT, SessionSubjectT, SessionT,
		ServiceMethodCallInputContextT]
}

// EntityMutatingResponseContext provides an abstraction for output contexts
// for mutating method calls.
type EntityMutatingResponseContext interface {
	EntityMutatingContext
	EntityMethodResponseContext
	ServiceMutatingMethodCallOutputContext
}

// EntityMutatingMessage abstracts entity mutating method requests and responses.
type EntityMutatingMessage interface {
	EntityMethodMessage
	ServiceMutatingMethodMessage

	EntityMutatingContext() EntityMutatingContext
}

// EntityMutatingRequest abstracts entity mutating requests.
type EntityMutatingRequest[
	SessionIDNumT SessionIDNum, SessionRefKeyT SessionRefKey[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalRefKeyT TerminalRefKey[TerminalIDNumT],
	UserIDNumT UserIDNum, UserRefKeyT UserRefKey[UserIDNumT],
	SessionSubjectT SessionSubject[
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
	],
	SessionT Session[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
		SessionSubjectT,
	],
	ServiceMethodCallInputContextT ServiceMethodCallInputContext[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
		SessionSubjectT,
		SessionT,
	],
	EntityMutatingRequestContextT EntityMutatingRequestContext[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT, SessionSubjectT, SessionT,
		ServiceMethodCallInputContextT],
] interface {
	EntityMutatingMessage
	ServiceMutatingMethodCallInput[
		SessionIDNumT, SessionRefKeyT, TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT, SessionSubjectT, SessionT,
		ServiceMethodCallInputContextT, EntityMutatingRequestContextT]

	EntityMutatingRequestContext() EntityMutatingRequestContextT
}

// EntityMutatingResponse abstracts entity mutating responses.
type EntityMutatingResponse interface {
	EntityMutatingMessage
	ServiceMutatingMethodCallOutput

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
