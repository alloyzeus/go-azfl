package azcore

import "time"

// ActionInfo holds information about an action.
type ActionInfo interface {
	// MethodCallID returns the ID of the method call this action initiated through.
	MethodCallID() MethodCallID

	// Actor returns the subject who executed the action. Must not be empty
	// in server, might be empty in clients, might be queryable.
	Actor() Subject

	// Session returns the session by the actor to perform the action.
	// Must not be empty in server, might be empty in clients.
	Session() Session

	// DelegationInfo returns the information about the delegation if this
	// action was delegated to other subject or session.
	DelegationInfo() DelegationInfo

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
type DelegationInfo interface {
	// ParentDelegationInfo returns the delegation parent of this delegation.
	ParentDelegationInfo() DelegationInfo
}
