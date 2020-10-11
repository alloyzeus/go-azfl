package azia

import "github.com/alloyzeus/go-azcore/azcore"

// Subject is an object which could be involved in an action.
type Subject interface {
	AZSubject()

	// TerminalRefKey returns the ref-key of the terminal for this subject.
	TerminalRefKey() TerminalRefKey

	// IsRepresentingAUser returns true if this subject is representing
	// a user, i.e., the application is a user-agent, not a service application.
	//
	// If this method returns true, UserRefKey must return a valid ref-key of
	// the user.
	IsRepresentingAUser() bool

	// UserRefKey returns the ref-key of the user this subject represents.
	UserRefKey() UserRefKey
}

type UserRefKey interface {
	azcore.RefKey
	AZUserRefKey()
}

type TerminalRefKey interface {
	azcore.RefKey
	AZTerminalRefKey()
}

type ActionInfo interface {
	// CallID returns the ID of the method call this action initiated through.
	CallID() azcore.MethodCallID

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

type DelegationInfo interface {
	DelegationInfo()
}

type Session interface {
	// SessionRefKey returns the identifier of this session.
	SessionRefKey() SessionRefKey

	// ParentSessionRefKey returns the ID of the session which was used to
	// create this session.
	ParentSessionRefKey() SessionRefKey

	// Subject returns the subject this session is for.
	Subject() Subject
}

type SessionRefKey interface {
	azcore.RefKey
	AZSessionRefKey()
}
