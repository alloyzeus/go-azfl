package azcore

// A Terminal is an object which could act within the system, i.e., an agent.
type Terminal interface {
	// RefKey returns the identifier of this Terminal instance.
	RefKey() TerminalRefKey

	// PrincipalUser returns the ref-key of the User, if any, who authorized
	// this instance of Terminal.
	PrincipalUser() UserRefKey
}

// TerminalID abstracts the identifiers of Terminal entity instances.
type TerminalID interface {
	EID
	AZTerminalID()
}

// TerminalRefKey is used to refer to a Terminal entity instance.
type TerminalRefKey interface {
	RefKey

	// TerminalID returns only the ID part of this ref-key.
	TerminalID() TerminalID
}
