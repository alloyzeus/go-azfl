package azcore

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
}

// DelegationInfo holds information about delegation for an action if that
// action was delegated.
type DelegationInfo interface {
	// ParentDelegationInfo returns the delegation parent of this delegation.
	ParentDelegationInfo() DelegationInfo
}
