package azcore

import "github.com/alloyzeus/go-azfl/azfl/azid"

// A Terminal is an object which could act within the system, i.e., an agent.
type Terminal[TerminalIDNumT TerminalIDNum, UserIDNumT UserIDNum] interface {
	// RefKey returns the identifier of this Terminal instance.
	RefKey() TerminalRefKey[TerminalIDNumT]

	// PrincipalUser returns the ref-key of the User, if any, who authorized
	// this instance of Terminal.
	PrincipalUser() UserRefKey[UserIDNumT]
}

type TerminalIDNumMethods interface {
	AZTerminalIDNum()
}

// TerminalIDNum abstracts the identifiers of Terminal entity instances.
type TerminalIDNum interface {
	azid.IDNum

	TerminalIDNumMethods
}

// TerminalRefKey is used to refer to a Terminal entity instance.
type TerminalRefKey[IDNumT TerminalIDNum] interface {
	azid.RefKey[IDNumT]

	// TerminalIDNum returns only the ID part of this ref-key.
	TerminalIDNum() IDNumT
}
