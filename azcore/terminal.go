package azcore

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
