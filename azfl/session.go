package azcore

import "github.com/alloyzeus/go-azfl/azfl/azid"

// Session represents information about a session. Every action can
// only be performed with an active session. A session is obtained through
// authorization, or authentication, of a Terminal.
//
//TODO: scope, expiry
type Session[
	SessionIDNumT SessionIDNum, SessionRefKeyT SessionRefKey[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalRefKeyT TerminalRefKey[TerminalIDNumT],
	UserIDNumT UserIDNum, UserRefKeyT UserRefKey[UserIDNumT],
	SubjectT Subject[TerminalIDNumT, UserIDNumT],
] interface {
	// RefKey returns the identifier of this Session instance.
	RefKey() SessionRefKeyT

	// ParentSessionRefKey returns the identifier of the session which
	// was used to create this session.
	ParentSessionRefKey() SessionRefKeyT

	// Subject returns the subject this session is for.
	Subject() SubjectT

	// IsTerminal returns true if the authorized terminal is the same as termRef.
	IsTerminal(termRef TerminalRefKeyT) bool

	// IsUserSubject returns true if the subject is a user instead of
	// a service application.
	IsUserSubject() bool

	// IsUser checks if this session is represeting a particular user.
	IsUser(userRef UserRefKeyT) bool
}

type SessionIDNumMethods interface {
	AZSessionIDNum()
}

// SessionIDNum abstracts the identifiers of Session entity instances.
type SessionIDNum interface {
	azid.IDNum

	SessionIDNumMethods
}

// SessionRefKey is used to refer to a Session entity instance.
type SessionRefKey[IDNumT SessionIDNum] interface {
	azid.RefKey[IDNumT]

	// SessionIDNum returns only the ID part of this ref-key.
	SessionIDNum() IDNumT
}
