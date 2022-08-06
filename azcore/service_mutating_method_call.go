package azcore

// ServiceMutatingMethodCallInputContext abstracts mutating method request contexts.
type ServiceMutatingMethodCallInputContext[
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
	ServiceMethodCallInputContextT ServiceMethodCallInputContext[
		SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT,
		SessionT, ServiceMethodIdempotencyKeyT],
	ServiceMethodIdempotencyKeyT ServiceMethodIdempotencyKey,
] interface {
	ServiceMutatingMethodContext
	ServiceMethodCallInputContext[
		SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT,
		UserIDNumT, UserIDT,
		SessionSubjectT,
		SessionT, ServiceMethodIdempotencyKeyT]
}

// ServiceMutatingMethodCallOutputContext abstracts mutating
// method response contexts.
type ServiceMutatingMethodCallOutputContext interface {
	ServiceMutatingMethodContext
	ServiceMethodCallOutputContext
}
