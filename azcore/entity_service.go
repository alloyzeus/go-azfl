package azcore

//region Method

// EntityMethodMessage abstracts the messages, i.e., requests and responses.
type EntityMethodMessage interface {
	ServiceMethodMessage

	EntityMethodContext() EntityMethodContext
}

// EntityMethodCallInput abstracts all entity method requests messages.
type EntityMethodCallInput[
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
	EntityMethodCallInputContextT EntityMethodCallInputContext[
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

	EntityMethodCallInputContext() EntityMethodCallInputContextT
}

// EntityMethodCallOutput abstracts all entity method response messages.
type EntityMethodCallOutput interface {
	ServiceMethodCallOutput
	EntityMethodMessage

	EntityMethodCallOutputContext() EntityMethodCallOutputContext
}

//endregion

//region Context

// EntityMethodContext provides an abstraction for all operations which
// apply to entity instances.
type EntityMethodContext interface {
	ServiceMethodContext
}

// EntityMethodCallInputContext is an abstraction for all method call
// input contexts.
type EntityMethodCallInputContext[
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

// EntityMethodCallOutputContext is an abstraction for all method call
// output contexts.
type EntityMethodCallOutputContext interface {
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

// EntityMutatingMethodCallContext provides an abstraction for input contexts
// for mutating method calls.
type EntityMutatingMethodCallContext[
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
	EntityMethodCallInputContext[
		SessionIDNumT, SessionRefKeyT, TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT, SessionSubjectT, SessionT]
	ServiceMutatingMethodCallInputContext[
		SessionIDNumT, SessionRefKeyT, TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT, SessionSubjectT, SessionT,
		ServiceMethodCallInputContextT]
}

// EntityMutatingMethodCallOutputContext provides an abstraction for output contexts
// for mutating method calls.
type EntityMutatingMethodCallOutputContext interface {
	EntityMutatingContext
	EntityMethodCallOutputContext
	ServiceMutatingMethodCallOutputContext
}

// EntityMutatingMessage abstracts entity mutating method requests and responses.
type EntityMutatingMessage interface {
	EntityMethodMessage
	ServiceMutatingMethodMessage

	EntityMutatingContext() EntityMutatingContext
}

// EntityMutatingCallInput abstracts entity mutating requests.
type EntityMutatingCallInput[
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
	EntityMutatingMethodCallContextT EntityMutatingMethodCallContext[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT, SessionSubjectT, SessionT,
		ServiceMethodCallInputContextT],
] interface {
	EntityMutatingMessage
	ServiceMutatingMethodCallInput[
		SessionIDNumT, SessionRefKeyT, TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT, SessionSubjectT, SessionT,
		ServiceMethodCallInputContextT, EntityMutatingMethodCallContextT]

	EntityMutatingMethodCallContext() EntityMutatingMethodCallContextT
}

// EntityMutatingMethodCallOutput abstracts entity mutating responses.
type EntityMutatingMethodCallOutput interface {
	EntityMutatingMessage
	ServiceMutatingMethodCallOutput

	EntityMutatingMethodCallOutputContext() EntityMutatingMethodCallOutputContext
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
