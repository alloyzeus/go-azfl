package azcore

import "time"

// OperationInfo holds information about an action.
type OperationInfo[
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
] interface {
	// MethodOpID returns the ID of the method call this action initiated through.
	MethodOpID() ServiceMethodIdempotencyKey

	// Actor returns the subject who executed the action. Must not be empty
	// in server, might be empty in clients, might be queryable.
	Actor() SessionSubjectT

	// Session returns the session by the actor to perform the action.
	// Must not be empty in server, might be empty in clients.
	Session() SessionT

	// DelegationInfo returns the information about the delegation if this
	// action was delegated to other subject or session.
	DelegationInfo() OperationDelegationInfo[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT, SessionSubjectT, SessionT]

	// Timestamp returns the time when the action made the effect. This should
	// be obtained from the lowest level, e.g., database or file system.
	//
	// An analogy: in ordering process in a restaurant, the
	// timestamp is the time when the cook declared that the food is ready
	// to be served.
	Timestamp() *time.Time

	// // BeginTime returns the time when the action was initiated.
	// BeginTime() *time.Time
	// // EndTime return the time when the action was ended.
	// EndTime() *time.Time
}

// OperationDelegationInfo holds information about delegation for an action
// if that action was delegated.
//
//TODO: actual info
type OperationDelegationInfo[
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
] interface {
	// ParentDelegationInfo returns the delegation parent of this delegation.
	ParentDelegationInfo() OperationDelegationInfo[
		SessionIDNumT, SessionRefKeyT,
		TerminalIDNumT, TerminalRefKeyT,
		UserIDNumT, UserRefKeyT, SessionSubjectT, SessionT]

	// Actor returns the subject who delegated the action. Must not be empty
	// in server, might be empty in clients, might be queryable.
	Actor() SessionSubjectT

	// Session returns the session by the actor to delegate the action.
	// Must not be empty in server, might be empty in clients.
	Session() SessionT
}
