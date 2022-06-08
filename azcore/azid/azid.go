// Package azid defines the encodings of entity identifiers and reference keys
// within AZ.
//
// There are two formats defined here: textual (azid-text) and
// binary (azid-bin).
package azid

// Error abstracts all errors emitted by this package.
type Error interface {
	error
}
