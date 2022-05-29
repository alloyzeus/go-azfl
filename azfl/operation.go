package azcore

import "time"

// OperationInfo holds information about an action.
type OperationInfo[
	SessionIDNumT SessionIDNum, TerminalIDNumT TerminalIDNum, UserIDNumT UserIDNum,
	SubjectT Subject[TerminalIDNumT, UserIDNumT],
	SessionT Session[
		SessionIDNumT, SessionRefKey[SessionIDNumT],
		TerminalIDNumT, TerminalRefKey[TerminalIDNumT],
		UserIDNumT, UserRefKey[UserIDNumT],
		SubjectT,
	],
] interface {
	// MethodOpID returns the ID of the method call this action initiated through.
	MethodOpID() ServiceMethodOpID

	// Actor returns the subject who executed the action. Must not be empty
	// in server, might be empty in clients, might be queryable.
	Actor() SubjectT

	// Session returns the session by the actor to perform the action.
	// Must not be empty in server, might be empty in clients.
	Session() SessionT

	// DelegationInfo returns the information about the delegation if this
	// action was delegated to other subject or session.
	DelegationInfo() DelegationInfo[
		SessionIDNumT, TerminalIDNumT, UserIDNumT, SubjectT, SessionT]

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

// DelegationInfo holds information about delegation for an action if that
// action was delegated.
//
//TODO: actual info
type DelegationInfo[
	SessionIDNumT SessionIDNum, TerminalIDNumT TerminalIDNum, UserIDNumT UserIDNum,
	SubjectT Subject[TerminalIDNumT, UserIDNumT],
	SessionT Session[
		SessionIDNumT, SessionRefKey[SessionIDNumT],
		TerminalIDNumT, TerminalRefKey[TerminalIDNumT],
		UserIDNumT, UserRefKey[UserIDNumT],
		SubjectT,
	],
] interface {
	// ParentDelegationInfo returns the delegation parent of this delegation.
	ParentDelegationInfo() DelegationInfo[
		SessionIDNumT, TerminalIDNumT, UserIDNumT, SubjectT, SessionT]

	// Actor returns the subject who delegated the action. Must not be empty
	// in server, might be empty in clients, might be queryable.
	Actor() SubjectT

	// Session returns the session by the actor to delegate the action.
	// Must not be empty in server, might be empty in clients.
	Session() SessionT
}
