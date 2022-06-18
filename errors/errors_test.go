package errors

import "testing"

const errNotFound = errorConstantDescriptor("not found")

func TestAssorted(t *testing.T) {
	var nf error = errNotFound
	var nf2 error = errNotFound
	var nf3 error = errorConstantDescriptor("not found")
	if nf != errorConstantDescriptor("not found") {
		t.Errorf("%#v %#v", nf, errorConstantDescriptor("not found"))
	}
	if nf != nf2 {
		t.Errorf("nf == nf2")
	}
	if nf != nf3 {
		t.Errorf("nf == nf3")
	}
	if _, ok := nf.(errorConstantDescriptor); !ok {
		t.Errorf("nf.(errorConstantDescriptor")
	}
	if nf3.Error() == "" {
		t.Errorf(`nf3.Error() == "" -- %q`, nf3.Error())
	}
	if IsEntNotFoundError(nf) {
		t.Errorf("IsEntNotFound(nf)")
	}
}
