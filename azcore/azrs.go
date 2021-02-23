package azcore

import (
	"github.com/rez-go/crock32"
)

// AZRS, which could be interpreted as AZ Reference String, is a rule for
// text-based encoding for identifiers within AZ framework. It limits the
// characters which can be used in encoded texts.
//
// AZRS is a limited implementation of the URL concept.

// AZRSObject provides an interface for objects which could be encoded
// to / decoded from AZRS encoding.
type AZRSObject interface {
	AZRSMarshalable
}

// AZRSMarshalable is an interface for objects which could be marshaled
// with AZRS encoding.
type AZRSMarshalable interface {
	AZRS() string
}

// An AZRSUnmarshalable is an object which could load AZRS-encoded
// string into itself.
type AZRSUnmarshalable interface {
	UnmarshalAZRS(s string) error
}

var (
	// AZRSEncode encodes azwire-encoded ref-key.
	AZRSEncode = crock32.Encode
	// AZRSDecode decodes azwire-encoded ref-key from a string.
	AZRSDecode = crock32.Decode
)
