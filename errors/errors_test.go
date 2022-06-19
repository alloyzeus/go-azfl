package errors

import "testing"

const errNotFound = constantErrorDescriptor("not found")

func TestAssorted(t *testing.T) {
	var nf error = errNotFound
	var nf2 error = errNotFound
	var nf3 error = constantErrorDescriptor("not found")
	if nf != constantErrorDescriptor("not found") {
		t.Errorf("%#v %#v", nf, constantErrorDescriptor("not found"))
	}
	if nf != nf2 {
		t.Errorf("nf == nf2")
	}
	if nf != nf3 {
		t.Errorf("nf == nf3")
	}
	if _, ok := nf.(constantErrorDescriptor); !ok {
		t.Errorf("nf.(errorConstantDescriptor")
	}
	if nf3.Error() == "" {
		t.Errorf(`nf3.Error() == "" -- %q`, nf3.Error())
	}
	if IsEntNotFoundError(nf) {
		t.Errorf("IsEntNotFound(nf)")
	}
}

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
