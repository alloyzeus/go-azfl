package errors

import (
	"reflect"
	"testing"
)

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
	if IsEntityNotFoundError(nf) {
		t.Errorf("IsEntNotFound(nf)")
	}
}

func assert(t *testing.T, expected interface{}, actual interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\tExpected: %#v\n\tActual: %#v", expected, actual)
	}
}

func assertNotEqual(t *testing.T, referenceValue interface{}, actualValue interface{}) {
	t.Helper()
	if reflect.DeepEqual(referenceValue, actualValue) {
		t.Fatalf("\n\tNot expected: %#v\n\tActual: %#v", referenceValue, actualValue)
	}
}
