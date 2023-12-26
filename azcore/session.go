package azcore

import "github.com/alloyzeus/go-azfl/v2/azid"

// Session represents information about a session. Every action can
// only be performed with an active session. A session is obtained through
// authorization, or authentication, of a Terminal.
//
// TODO: scope, expiry.
type Session[
	SessionIDNumT SessionIDNum, SessionIDT SessionID[SessionIDNumT],
	TerminalIDNumT TerminalIDNum, TerminalIDT TerminalID[TerminalIDNumT],
	UserIDNumT UserIDNum, UserIDT UserID[UserIDNumT],
	SessionSubjectT SessionSubject[
		TerminalIDNumT, TerminalIDT, UserIDNumT, UserIDT,
	],
] interface {
	// ID returns the identifier of this Session instance.
	ID() SessionIDT

	// DelegateSession returns the session that is a delegate of
	// this session. It returns nil if this session is not a delegated session.
	//
	// A delegation is commonly used when a service is accessing another
	// service on the behalf of a user.
	//
	//     User --> Service A --> Service B
	//
	// In this example, the user is delegating Service A to accces their
	// data in Service B. As seen by Service B, the subject of the Session
	// is the delegating user and this method will return the Session of
	// Service A.
	DelegateSession() Session[SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT, UserIDNumT, UserIDT, SessionSubjectT]

	// ImpersonatorSession returns the session that is impersonating
	// the subject of this session, i.e., the session that was used to create
	// this session. It returns nil if this session is not an impersonation.
	//
	//     User A (Admin) --> (User B) --> Service
	//
	// In this example, User A, who is an admin, impersonate User B to access
	// the Service. The subject of the Session as seen by the Service is
	// User B and this method will return the Session of User A.
	ImpersonatorSession() Session[SessionIDNumT, SessionIDT,
		TerminalIDNumT, TerminalIDT, UserIDNumT, UserIDT, SessionSubjectT]

	// Subject returns the subject of this session.
	Subject() SessionSubjectT

	// IsTerminal returns true if the authorized terminal is the same as termRef.
	IsTerminal(termRef TerminalIDT) bool

	// HasUserAsSubject returns true if the subject is a user instead of
	// a service application.
	HasUserAsSubject() bool

	// IsUser checks if this session is representing a particular user.
	IsUser(userRef UserIDT) bool
}

type SessionIDNumMethods interface {
}

// SessionIDNum abstracts the identifiers of Session entity instances.
type SessionIDNum interface {
	azid.IDNum

	SessionIDNumMethods
}

// SessionID is used to refer to a Session entity instance.
type SessionID[IDNumT SessionIDNum] interface {
	azid.ID[IDNumT]

	// SessionIDNum returns only the ID part of this ref-key.
	SessionIDNum() IDNumT
}
