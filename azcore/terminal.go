package azcore

import (
	"context"

	"github.com/alloyzeus/go-azfl/v2/azid"
)

// A Terminal is an object which could act within the system, i.e., an agent.
type Terminal[
	TerminalIDNumT TerminalIDNum, TerminalIDT TerminalID[TerminalIDNumT],
	UserIDNumT UserIDNum, UserIDT UserID[UserIDNumT],
] interface {
	// ID returns the identifier of this Terminal instance.
	ID() TerminalIDT

	// PrincipalUser returns the ref-key of the User, if any, who authorized
	// this instance of Terminal.
	PrincipalUser() UserIDT
}

type TerminalIDNumMethods interface {
}

// TerminalIDNum abstracts the identifiers of Terminal entity instances.
type TerminalIDNum interface {
	azid.IDNum

	TerminalIDNumMethods
}

// TerminalID is used to refer to a Terminal entity instance.
type TerminalID[IDNumT TerminalIDNum] interface {
	azid.ID[IDNumT]

	// TerminalIDNum returns only the ID part of this ref-key.
	TerminalIDNum() IDNumT
}

func NewTerminalRef[
	TerminalIDNumT TerminalIDNum, TerminalIDT TerminalID[TerminalIDNumT],
	UserIDNumT UserIDNum, UserIDT UserID[UserIDNumT],
	TerminalT Terminal[TerminalIDNumT, TerminalIDT, UserIDNumT, UserIDT],
](
	id TerminalIDT,
	instanceResolver func(context.Context, TerminalIDT) (TerminalT, error),
) TerminalRef[TerminalIDNumT, TerminalIDT, UserIDNumT, UserIDT, TerminalT] {
	return TerminalRef[TerminalIDNumT, TerminalIDT, UserIDNumT, UserIDT, TerminalT]{
		id:               id,
		instanceResolver: instanceResolver,
	}
}

type TerminalRef[
	TerminalIDNumT TerminalIDNum, TerminalIDT TerminalID[TerminalIDNumT],
	UserIDNumT UserIDNum, UserIDT UserID[UserIDNumT],
	TerminalT Terminal[TerminalIDNumT, TerminalIDT, UserIDNumT, UserIDT],
] struct {
	id               TerminalIDT
	instanceResolver func(context.Context, TerminalIDT) (TerminalT, error)
}
