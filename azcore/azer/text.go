package azer

import "github.com/rez-go/crock32"

// TextMarshalable provides an interface definition for objects which able to
// provide an azer-text-encoded representation of itself.
type TextMarshalable interface {
	AZERText() string
}

var (
	// TextEncode encodes azer-bin-encoded ref-key.
	TextEncode = crock32.EncodeLower
)

// TextDecode decodes azer-bin-encoded ref-key from a string.
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
		return nil, err //TODO: wrap/translate
	}
	if len(csEncoded) == 2 {
		//TODO: checksum
	}
	return dataBytes, err
}
