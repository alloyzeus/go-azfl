package azcore

import "github.com/alloyzeus/go-azfl/azfl/azid"

// Session represents information about a session. Every action can
// only be performed with an active session. A session is obtained through
// authorization, or authentication, of a Terminal.
type Session[
	SessionIDNumT SessionIDNum, TerminalIDNumT TerminalIDNum, UserIDNumT UserIDNum,
] interface {
	// RefKey returns the identifier of this Session instance.
	RefKey() SessionRefKey[SessionIDNumT]

	// ParentSessionRefKey returns the identifier of the session which
	// was used to create this session.
	ParentSessionRefKey() SessionRefKey[SessionIDNumT]

	// Subject returns the subject this session is for.
	Subject() Subject[TerminalIDNumT, UserIDNumT]
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
