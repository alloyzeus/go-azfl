package azid

import (
	"github.com/rez-go/crock32"

	"github.com/alloyzeus/go-azfl/azfl/errors"
)

// TextMarshalable is an interface definition for objects which able to
// provide an azid-text representation of itself.
type TextMarshalable interface {
	AZIDText() string
}

// A TextUnmarshalable is an object which able to load an
// azid-text representation into itself.
type TextUnmarshalable interface {
	UnmarshalAZIDText(s string) error
}

var (
	// TextEncode encodes azid-bin ref-key.
	TextEncode = crock32.EncodeLower
)

// TextDecode decodes azid-bin ref-key from a string.
func TextDecode(s string) ([]byte, error) {
	var dataEncoded []byte
	var csEncoded []byte
	inAsBytes := []byte(s)
	for i, c := range inAsBytes {
		if c == 'U' || c == 'u' {
			if i+1 < len(inAsBytes) {
				if c = inAsBytes[i+1]; c == 'N' || c == 'n' {
					dataEncoded = inAsBytes[:i]
					csEncoded = inAsBytes[i+2 : i+4]
					break
				}
			} else {
				dataEncoded = inAsBytes[:i]
				break
			}
		}
	}
	if dataEncoded == nil {
		dataEncoded = []byte(s)
	}
	dataBytes, err := crock32.Decode(string(dataEncoded))
	if err != nil {
		return nil, errors.ArgWrap("", "data decoding", err)
	}
	if len(csEncoded) == 2 {
		//TODO: checksum
	}
	return dataBytes, nil
}
