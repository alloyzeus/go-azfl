package errors

const (
	ErrPlaceholder       = constantErrorDescriptor(placeholderErrorText)
	placeholderErrorText = `` +
		`This is a placeholder error. It means that the developer of the ` +
		`routine where you get the error from has not properly thought about ` +
		`what kind of errors should be returned from the routine. You could ` +
		`treat this error as 'internal error' or such, but if you really need ` +
		`to be able to react based on the different kind of errors, do not ` +
		`depend on this error, instead contact the developer of the routine ` +
		`where this error was returned from. Provide your use cases of the ` +
		`errors to them.`
)
