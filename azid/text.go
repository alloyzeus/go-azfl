package azid

import (
	"github.com/rez-go/crock32"

	errors "github.com/alloyzeus/go-azfl/v2/azerrs"
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
	var checksumEncoded []byte
	inputAsBytes := []byte(s)
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
		dataEncoded = []byte(s)
	}
	dataBytes, err := crock32.Decode(string(dataEncoded))
	if err != nil {
		return nil, errors.Arg1().Desc(errors.ErrValueMalformed).Wrap(err)
	}
	if len(checksumEncoded) == 2 {
		//TODO: checksum
	}
	return dataBytes, nil
}
