package azcore

// ServiceMutatingMethodCallInputContext abstracts mutating method request contexts.
type ServiceMutatingMethodCallInputContext[
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
		SessionT, ServiceMethodIdempotencyKeyT],
	ServiceMethodIdempotencyKeyT ServiceMethodIdempotencyKey,
] interface {
	ServiceMutatingMethodContext
	ServiceMethodCallInputContext[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT,
		SessionSubjectT,
		SessionT, ServiceMethodIdempotencyKeyT]
}

// ServiceMutatingMethodCallOutputContext abstracts mutating
// method response contexts.
type ServiceMutatingMethodCallOutputContext interface {
	ServiceMutatingMethodContext
	ServiceMethodCallOutputContext
}
