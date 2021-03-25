package azcore

import "github.com/alloyzeus/go-azfl/azfl/azid"

// A Terminal is an object which could act within the system, i.e., an agent.
type Terminal interface {
	// RefKey returns the identifier of this Terminal instance.
	RefKey() TerminalRefKey

	// PrincipalUser returns the ref-key of the User, if any, who authorized
	// this instance of Terminal.
	PrincipalUser() UserRefKey
}

// TerminalIDNum abstracts the identifiers of Terminal entity instances.
type TerminalIDNum interface {
	azid.IDNum

	AZTerminalIDNum()
}

// TerminalRefKey is used to refer to a Terminal entity instance.
type TerminalRefKey interface {
	azid.RefKey

	// TerminalIDNum returns only the ID part of this ref-key.
	TerminalIDNum() TerminalIDNum
}
