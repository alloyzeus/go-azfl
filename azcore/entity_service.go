package azcore

//region Method

// EntityOpMessage abstracts the messages, i.e., requests and responses.
type EntityOpMessage interface {
	ServiceOpMessage

	EntityOpContext() EntityOpContext
}

//endregion

//region Context

// EntityOpContext provides an abstraction for all operations which
// apply to entity instances.
type EntityOpContext interface {
	ServiceOpContext
}

// EntityOpCallContext is an abstraction for all method call
// input contexts.
type EntityOpCallContext[
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
	ServiceOpIdempotencyKeyT ServiceOpIdempotencyKey,
] interface {
	EntityOpContext
	ServiceOpCallContext[
		SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT,
		SessionT, ServiceOpIdempotencyKeyT,
	]
}

// EntityOpResultContext is an abstraction for all method call
// output contexts.
type EntityOpResultContext interface {
	EntityOpContext
	ServiceOpResultContext
}

//endregion

//region MutatingContext

// EntityMutationOpContext is a specialization of EntityOperationContext which
// is used for operations which make any change to the entity.
type EntityMutationOpContext interface {
	EntityOpContext
	ServiceMutationOpContext
}

// EntityMutatingOpCallContext provides an abstraction for input contexts
// for mutating method calls.
type EntityMutatingOpCallContext[
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
	ServiceOpCallContextT ServiceOpCallContext[
		SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT,
		SessionT, ServiceOpIdempotencyKeyT],
	ServiceOpIdempotencyKeyT ServiceOpIdempotencyKey,
] interface {
	EntityMutationOpContext
	EntityOpCallContext[
		SessionIDNumT, SessionIDT, TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT, SessionSubjectT, SessionT,
		ServiceOpIdempotencyKeyT]
	ServiceMutationOpCallContext[
		SessionIDNumT, SessionIDT, TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT, SessionSubjectT, SessionT,
		ServiceOpCallContextT, ServiceOpIdempotencyKeyT]
}

// EntityMutatingOpResultContext provides an abstraction for output contexts
// for mutating method calls.
type EntityMutatingOpResultContext interface {
	EntityMutationOpContext
	EntityOpResultContext
	ServiceMutationOpResultContext
}

// EntityMutatingMessage abstracts entity mutating method requests and responses.
type EntityMutatingMessage interface {
	EntityOpMessage
	ServiceMutationOpMessage

	EntityMutatingContext() EntityMutationOpContext
}

//endregion

//region Service

// EntityService provides an abstraction for all entity services. This
// abstraction is used by both client and server.
type EntityService interface {
}

//endregion

//region ServiceBase

// EntityServiceBase provides a basic implementation for EntityService.
type EntityServiceBase struct{}

var _ EntityService = &EntityServiceBase{}

//endregion

//region ServiceClient

// EntityServiceClient provides an abstraction for all entity service clients.
type EntityServiceClient interface {
	EntityService
}

//endregion

//region ServiceServer

// EntityServiceServer provides an abstraction for all entity service servers.
type EntityServiceServer interface {
	EntityService
}

//endregion
