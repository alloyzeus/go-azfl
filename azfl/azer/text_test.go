package azer_test

import (
	"bytes"
	"testing"

	"github.com/alloyzeus/go-azfl/azfl/azer"
)

var _ azer.TextMarshalable = adjunctRefKey{}

func (refKey adjunctRefKey) AZERText() string {
	bin := refKey.AZERBin()
	return azer.TextEncode(bin)
}

func TestAZERTextAdjunctRefKeyEncode(t *testing.T) {
	testCases := []struct {
		in  adjunctRefKey
		out string
	}{
		{adjunctRefKey{}, "801164g000000000"},
		{adjunctRefKey{0, 1}, "801164g000000001"},
		{adjunctRefKey{1, 0}, "801164g000002000"},
		{adjunctRefKey{1, 1}, "801164g000002001"},
	}

	for _, testCase := range testCases {
		s := testCase.in.AZERText()
		if s != testCase.out {
			t.Errorf("Expected: %#v, got: %#v", testCase.out, s)
		}
	}
}

func TestAZERTextDecode(t *testing.T) {
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
		r, err := azer.TextDecode(testCase.in)
		if err != testCase.err {
			t.Errorf("Expected: %#v, got: %#v", testCase.err, err)
		}
		if !bytes.Equal(r, testCase.out) {
			t.Errorf("Expected: %#v, got: %#v", testCase.out, r)
		}
	}
}
