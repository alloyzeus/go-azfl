package azid

import (
	"github.com/rez-go/crock32"

	"github.com/alloyzeus/go-azfl/v2/errors"
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
func TextDecode(input string) ([]byte, error) {
	var dataEncoded []byte
	var checksumEncoded []byte
	inputAsBytes := []byte(input)
	for i, c := range inputAsBytes {
		if c == 'U' || c == 'u' {
			if i+1 < len(inputAsBytes) {
				if c = inputAsBytes[i+1]; c == 'N' || c == 'n' {
					dataEncoded = inputAsBytes[:i]
					checksumEncoded = inputAsBytes[i+2 : i+4]
					break
				}
			} else {
				dataEncoded = inputAsBytes[:i]
				break
			}
		}
	}
	if dataEncoded == nil {
		dataEncoded = []byte(input)
	}
	dataBytes, err := crock32.Decode(string(dataEncoded))
	if err != nil {
		return nil, errors.ArgDW("input", errors.ErrValueMalformed, err)
	}
	if len(checksumEncoded) == 2 {
		//TODO: checksum
	}
	return dataBytes, nil
}
