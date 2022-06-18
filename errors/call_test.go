package errors

import "testing"

func TestCallErrorSetEmpty(t *testing.T) {
	var err error = &callErrorSet{}
	if err.Error() != "call error" {
		t.Errorf(`err.Error() != "call error" -- %q`, err.Error())
	}
}

func TestCallErrorSetContexUnspecified(t *testing.T) {
	var err error = &callErrorSet{callErrors: []CallError{ContextUnspecified()}}
	if err.Error() != "call: context unspecified" {
		t.Errorf(`err.Error() != "call: context unspecified" -- %q`, err.Error())
	}
}
