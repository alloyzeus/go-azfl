package azcore

// ServiceMutationOpCallContext abstracts mutating method request contexts.
type ServiceMutationOpCallContext[
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
	ServiceOpCallContextT ServiceOpCallContext[
		SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT,
		SessionT, ServiceOpIdempotencyKeyT],
	ServiceOpIdempotencyKeyT ServiceOpIdempotencyKey,
] interface {
	ServiceMutationOpContext
	ServiceOpCallContext[
		SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT,
		SessionT, ServiceOpIdempotencyKeyT]
}

// ServiceMutationOpResultContext abstracts mutating
// method response contexts.
type ServiceMutationOpResultContext interface {
	ServiceMutationOpContext
	ServiceOpResultContext
}
