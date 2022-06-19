package errors

import "testing"

func TestUnwrapCallErrorSetNil(t *testing.T) {
	var err error
	if callErrors := UnwrapCallErrorSet(err); len(callErrors) != 0 {
		t.Error("len(callErrors) != 0")
	}
}

func TestUnwrapCallErrorSetWrongType(t *testing.T) {
	var err error = ErrValueMalformed
	if callErrors := UnwrapCallErrorSet(err); len(callErrors) != 0 {
		t.Error("len(callErrors) != 0")
	}
}

func TestCallErrorSetEmpty(t *testing.T) {
	var err error = callErrorSet{}
	if err.Error() != "call error" {
		t.Errorf(`err.Error() != "call error" -- %q`, err.Error())
	}
	errSet := asErrorSet(err)
	if errSet == nil {
		t.Error("errSet == nil")
	}
	errs := errSet.Errors()
	if len(errs) != 0 {
		t.Error("len(errors) != 0")
	}
	callErrSet := asCallErrorSet(err)
	if callErrSet == nil {
		t.Error("callErrSet == nil")
	}
	callErrors := callErrSet.CallErrors()
	if len(callErrors) != 0 {
		t.Error("len(callErrors) != 0")
	}
	callErrors = UnwrapCallErrorSet(err)
	if len(callErrors) != 0 {
		t.Error("len(callErrors) != 0")
	}
}

func TestCallErrorSetContexUnspecified(t *testing.T) {
	var err error = CallSet(ContextUnspecified())
	if err.Error() != "call: context unspecified" {
		t.Errorf(`err.Error() != "call: context unspecified" -- %q`, err.Error())
	}
	errSet := asErrorSet(err)
	if errSet == nil {
		t.Error("errSet == nil")
	}
	errs := errSet.Errors()
	if len(errs) != 1 {
		t.Error("len(errors) != 1")
	}
	callErrSet := asCallErrorSet(err)
	if callErrSet == nil {
		t.Error("callErrSet == nil")
	}
	callErrors := callErrSet.CallErrors()
	if len(callErrors) != 1 {
		t.Error("len(callErrors) != 1")
	}
	callErrors = UnwrapCallErrorSet(err)
	if len(callErrors) != 1 {
		t.Error("len(callErrors) != 1")
	}
}
