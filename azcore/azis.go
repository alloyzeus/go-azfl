package azcore

// AZIS, which could be interpreted as AZ Identifier String, is a rule for
// text-based encoding for identifiers within AZ framework. It limits the
// characters which can be used in encoded texts.

// AZISObject provides an interface for objects which could be encoded
// to / decoded from AZIS encoding.
type AZISObject interface {
	AZISMarshalable
}

// AZISMarshalable is an interface for objects which could be marshaled
// with AZIS encoding.
type AZISMarshalable interface {
	AZIS() string
}

// An AZISUnmarshalable is an object which could load AZIS-encoded
// string into itself.
type AZISUnmarshalable interface {
	UnmarshalAZIS(s string) error
}
