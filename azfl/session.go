package azcore

// Session represents information about a session. Every action can
// only be performed with an active session. A session is obtained through
// authorization, or authentication, of a Terminal.
type Session interface {
	// RefKey returns the identifier of this Session instance.
	RefKey() SessionRefKey

	// ParentSessionRefKey returns the identifier of the session which
	// was used to create this session.
	ParentSessionRefKey() SessionRefKey

	// Subject returns the subject this session is for.
	Subject() Subject
}

// SessionID abstracts the identifiers of Session entity instances.
type SessionID interface {
	IDNum
	AZSessionID()
}

// SessionRefKey is used to refer to a Session entity instance.
type SessionRefKey interface {
	RefKey

	// SessionID returns only the ID part of this ref-key.
	SessionID() SessionID
}
