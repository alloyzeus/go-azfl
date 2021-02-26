// Package azer defines the encodings of entity reference within AZ.
//
// There are two formats defined here: textual (azer-text) and
// binary (azer-bin).
package azer

// Error abstracts all errors emitted by this package.
type Error interface {
	error
}
