package azcore

// SessionSubject is an object which could be involved in an action.
type SessionSubject[
	TerminalIDNumT TerminalIDNum, TerminalIDT TerminalID[TerminalIDNumT],
	UserIDNumT UserIDNum, UserIDT UserID[UserIDNumT],
] interface {
	AZSessionSubject()

	// TerminalID returns the ref-key of the terminal for this subject.
	TerminalID() TerminalIDT

	// IsRepresentingAUser returns true if this subject is representing
	// a user, i.e., the application is a user-agent, not a service application.
	//
	// If this method returns true, UserID must return a valid ref-key of
	// the user.
	IsRepresentingAUser() bool

	// UserID returns the ref-key of the user this subject represents.
	UserID() UserIDT
}
