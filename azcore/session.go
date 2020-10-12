package azcore

// Session represents information about a session. Every action can
// only be performed with an active session. A session is obtained through
// authorization, or authentication, of a Terminal.
type Session interface {
	// SessionRefKey returns the identifier of this session.
	SessionRefKey() SessionRefKey

	// ParentSessionRefKey returns the identifier of the session which
	// was used to create this session.
	ParentSessionRefKey() SessionRefKey

	// Subject returns the subject this session is for.
	Subject() Subject
}

// A SessionRefKey is a reference to a Session.
type SessionRefKey interface {
	RefKey
	AZSessionRefKey()
}
