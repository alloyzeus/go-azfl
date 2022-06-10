package azid_test

import (
	"bytes"
	"testing"

	"github.com/alloyzeus/go-azfl/azid"
)

var _ azid.TextMarshalable = adjunctID{}

func (id adjunctID) AZIDText() string {
	bin := id.AZIDBin()
	return azid.TextEncode(bin)
}

func TestAZIDTextAdjunctIDEncode(t *testing.T) {
	testCases := []struct {
		in  adjunctID
		out string
	}{
		{adjunctID{}, "801164g000000000"},
		{adjunctID{0, 1}, "801164g000000001"},
		{adjunctID{1, 0}, "801164g000002000"},
		{adjunctID{1, 1}, "801164g000002001"},
	}

	for _, testCase := range testCases {
		s := testCase.in.AZIDText()
		if s != testCase.out {
			t.Errorf("Expected: %#v, got: %#v", testCase.out, s)
		}
	}
}

func TestAZIDTextDecode(t *testing.T) {
	testCases := []struct {
		in  string
		out []byte
		err error
	}{
		{"ht0f", []byte{0x08, 0xe8, 0x0f}, nil},
		{"ht0fuNC0", []byte{0x08, 0xe8, 0x0f}, nil},
		{"ht0fUNC0", []byte{0x08, 0xe8, 0x0f}, nil},
		{"ht0fUnC0", []byte{0x08, 0xe8, 0x0f}, nil},
		{"ht0funC0", []byte{0x08, 0xe8, 0x0f}, nil},
		{"8T2J81007", []byte{0x08, 0xd0, 0xa4, 0x80, 0x80, 0x07}, nil},
		{"8T2J81007uN00", []byte{0x08, 0xd0, 0xa4, 0x80, 0x80, 0x07}, nil},
		{"8t2j81007", []byte{0x08, 0xd0, 0xa4, 0x80, 0x80, 0x07}, nil},
		{"8t2j81007Un00", []byte{0x08, 0xd0, 0xa4, 0x80, 0x80, 0x07}, nil},
	}

	for _, testCase := range testCases {
		r, err := azid.TextDecode(testCase.in)
		if err != testCase.err {
			t.Errorf("Expected: %#v, got: %#v", testCase.err, err)
		}
		if !bytes.Equal(r, testCase.out) {
			t.Errorf("Expected: %#v, got: %#v", testCase.out, r)
		}
	}
}
