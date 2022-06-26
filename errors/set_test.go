package errors

import "testing"

func TestUnwrapErrorSetNil(t *testing.T) {
	var err error
	if errs := UnwrapErrorSet(err); len(errs) != 0 {
		t.Error("len(errs) != 0")
	}
}

func TestUnwrapErrorSetWrongType(t *testing.T) {
	var err error = ErrValueMalformed
	if errs := UnwrapErrorSet(err); len(errs) != 0 {
		t.Error("len(errs) != 0")
	}
}
