package azcore

// SessionSubject is an object which could be involved in an action.
type SessionSubject[
	TerminalIDNumT TerminalIDNum, TerminalRefKeyT TerminalRefKey[TerminalIDNumT],
	UserIDNumT UserIDNum, UserRefKeyT UserRefKey[UserIDNumT],
] interface {
	AZSessionSubject()

	// TerminalRefKey returns the ref-key of the terminal for this subject.
	TerminalRefKey() TerminalRefKeyT

	// IsRepresentingAUser returns true if this subject is representing
	// a user, i.e., the application is a user-agent, not a service application.
	//
	// If this method returns true, UserRefKey must return a valid ref-key of
	// the user.
	IsRepresentingAUser() bool

	// UserRefKey returns the ref-key of the user this subject represents.
	UserRefKey() UserRefKeyT
}
