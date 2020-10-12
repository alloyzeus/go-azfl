package azcore

// TerminalRefKey represents a reference to a Terminal.
type TerminalRefKey interface {
	RefKey
	AZTerminalRefKey()
}
