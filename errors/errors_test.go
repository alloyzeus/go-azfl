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
	if IsEntNotFound(nf) {
		t.Errorf("IsEntNotFound(nf)")
	}
}

func TestWrapEmpty(t *testing.T) {
	var err error = Wrap("", nil)
	if u, ok := err.(Unwrappable); !ok {
		t.Errorf("err.(Unwrappable)")
	} else {
		inner := u.Unwrap()
		if inner != nil {
			t.Errorf("inner != nil")
		}
		if inner == DataErrUnspecified {
			t.Errorf("inner == DataErrUnspecified")
		}
	}

	inner := Unwrap(err)
	if inner != nil {
		t.Errorf("inner != nil")
	}
	if inner == DataErrUnspecified {
		t.Errorf("inner == DataErrUnspecified")
	}

	if err.Error() != "" {
		t.Errorf(`err.Error() != "" -- %q`, err.Error())
	}
}

func TestWrapNoContext(t *testing.T) {
	var err error = Wrap("", DataErrUnspecified)
	if u, ok := err.(Unwrappable); !ok {
		t.Errorf("err.(Unwrappable)")
	} else {
		inner := u.Unwrap()
		if inner == nil {
			t.Errorf("inner == nil")
		}
		if inner != DataErrUnspecified {
			t.Errorf("inner != DataErrUnspecified")
		}
	}

	inner := Unwrap(err)
	if inner == nil {
		t.Errorf("inner == nil")
	}
	if inner != DataErrUnspecified {
		t.Errorf("inner != DataErrUnspecified")
	}

	if err.Error() != "unspecified" {
		t.Errorf(`err.Error() != "unspecified" -- %q`, err.Error())
	}
}

func TestWrapNoErr(t *testing.T) {
	var err error = Wrap("unknown", nil)
	if u, ok := err.(Unwrappable); !ok {
		t.Errorf("err.(Unwrappable)")
	} else {
		inner := u.Unwrap()
		if inner != nil {
			t.Errorf("inner != nil")
		}
		if inner == DataErrUnspecified {
			t.Errorf("inner == DataErrUnspecified")
		}
	}

	inner := Unwrap(err)
	if inner != nil {
		t.Errorf("inner != nil")
	}
	if inner == DataErrUnspecified {
		t.Errorf("inner == DataErrUnspecified")
	}

	if err.Error() != "unknown" {
		t.Errorf(`err.Error() != "unknown" -- %q`, err.Error())
	}
}

func TestWrap(t *testing.T) {
	var err error = Wrap("launch code", DataErrUnspecified)
	if u, ok := err.(Unwrappable); !ok {
		t.Errorf("err.(Unwrappable)")
	} else {
		inner := u.Unwrap()
		if inner == nil {
			t.Errorf("inner == nil")
		}
		if inner != DataErrUnspecified {
			t.Errorf("inner != DataErrUnspecified")
		}
	}

	inner := Unwrap(err)
	if inner == nil {
		t.Errorf("inner == nil")
	}
	if inner != DataErrUnspecified {
		t.Errorf("inner != DataErrUnspecified")
	}

	if err.Error() != "launch code: unspecified" {
		t.Errorf(`err.Error() != "launch code: unspecified" -- %q`, err.Error())
	}
}

func TestDataMalformedBare(t *testing.T) {
	var err error = DataErrMalformed
	if err != DataErrMalformed {
		t.Error("err != DataErrMalformed")
	}
	if !IsDataMalformed(err) {
		t.Errorf("!IsDataMalformed(err)")
	}
	if err.Error() != "malformed" {
		t.Errorf(`err.Error() != "malformed" -- %q`, err.Error())
	}
}

func TestDataMalformedComplexNil(t *testing.T) {
	var err error = DataMalformed(nil)
	if !IsDataMalformed(err) {
		t.Error("!IsDataMalformed(err)")
	}
	if err.Error() != "malformed" {
		t.Errorf(`err.Error() != "malformed" -- %q`, err.Error())
	}
	if Unwrap(err) != nil {
		t.Errorf("Unwrap(err) == nil -- %#v", Unwrap(err))
	}
}

func TestDataMalformedComplexWithDetails(t *testing.T) {
	var err error = DataMalformed(DataErrInvalid)
	if !IsDataMalformed(err) {
		t.Error("!IsDataMalformed(err)")
	}
	if err.Error() != "malformed: invalid" {
		t.Errorf(`err.Error() != "malformed: invalid" -- %q`, err.Error())
	}
	if Unwrap(err) == nil {
		t.Errorf("Unwrap(err) == nil -- %#v", Unwrap(err))
	}
}

func TestEntityBlank(t *testing.T) {
	var err error = Ent("", nil)
	if err.Error() != "entity invalid" {
		t.Errorf(`err.Error() != "entity invalid" -- %q`, err.Error())
	}
	if entErr, ok := err.(EntityError); !ok {
		t.Error("err.(EntityError)")
	} else {
		if entErr.EntityIdentifier() != "" {
			t.Error(`entErr.EntityIdentifier() != ""`)
		}
	}
	if inner := Unwrap(err); inner != nil {
		t.Error("inner != nil")
	}
}

func TestEntityNoID(t *testing.T) {
	var err error = Ent("", DataErrMalformed)
	if err.Error() != "entity malformed" {
		t.Errorf(`err.Error() != "entity malformed" -- %q`, err.Error())
	}
	if entErr, ok := err.(EntityError); !ok {
		t.Error("err.(EntityError)")
	} else {
		if entErr.EntityIdentifier() != "" {
			t.Error(`entErr.EntityIdentifier() != ""`)
		}
	}
}

func TestEntityWithFields(t *testing.T) {
	var err error = EntFields(
		"user",
		Ent("name", DataErrEmpty),
		Ent("age", DataErrUnspecified))
	if err.Error() != "user invalid: name: empty, age: unspecified" {
		t.Errorf(`err.Error() != "user invalid: name: empty, age: unspecified" -- %q`, err.Error())
	}
}

func TestEntNotFound(t *testing.T) {
	var fooNotFound error = EntNotFound("foo", nil)
	if !IsEntNotFound(fooNotFound) {
		t.Errorf("!IsEntNotFound(fooNotFound)")
	}
	if fooNotFound.Error() != "foo: not found" {
		t.Errorf(`fooNotFound.Error() != "foo: not found" -- %q`, fooNotFound.Error())
	}
	var notFoundBare error = EntErrNotFound
	if !IsEntNotFound(notFoundBare) {
		t.Errorf("!IsEntNotFound(notFoundBare)")
	}
	if IsArgumentError(fooNotFound) {
		t.Errorf("IsArgumentError(fooNotFound)")
	}
	if IsCallError(fooNotFound) {
		t.Errorf("IsCallError(fooNotFound)")
	}
}

func TestEntInvalid(t *testing.T) {
	var fooInvalid error = EntInvalid("foo", nil)
	if !IsEntInvalid(fooInvalid) {
		t.Errorf("!IsEntInvalid(fooInvalid)")
	}
	if fooInvalid.Error() != "foo: invalid" {
		t.Errorf(`fooInvalid.Error() != "foo: invalid" -- %q`, fooInvalid.Error())
	}
	var notFoundBare error = DataErrInvalid
	if !IsEntInvalid(notFoundBare) {
		t.Errorf("!IsEntInvalid(notFoundBare)")
	}
	if IsArgumentError(fooInvalid) {
		t.Errorf("IsArgumentError(fooInvalid)")
	}
	if IsCallError(fooInvalid) {
		t.Errorf("IsCallError(fooInvalid)")
	}
}

func TestArgUnspecifiedEmpty(t *testing.T) {
	var err error = ArgUnspecified("")
	if err.Error() != "arg unspecified" {
		t.Errorf(`err.Error() != "arg unspecified" -- %q`, err.Error())
	}
	if argErr, ok := err.(ArgumentError); !ok {
		t.Error("err.(ArgumentError)")
	} else {
		if argErr == nil {
			t.Error("argErr == nil")
		}
		if argErr.ArgumentName() != "" {
			t.Errorf(`argErr.ArgumentName() == "" -- %q`, argErr.ArgumentName())
		}
	}
}

func TestArgUnspecifiedFoo(t *testing.T) {
	var err error = ArgUnspecified("foo")
	if err.Error() != "arg foo: unspecified" {
		t.Errorf(`err.Error() != "arg foo: unspecified" -- %q`, err.Error())
	}
	if argErr, ok := err.(ArgumentError); !ok {
		t.Error("err.(ArgumentError)")
	} else {
		if argErr == nil {
			t.Error("argErr == nil")
		}
		if argErr.ArgumentName() != "foo" {
			t.Errorf(`argErr.ArgumentName() == "foo" -- %q`, argErr.ArgumentName())
		}
	}
	inner := Unwrap(err)
	if inner == nil {
		t.Error("inner == nil")
	}
	if inner != DataErrUnspecified {
		t.Error("inner != DataErrUnspecified")
	}
}

func TestArgEmpty(t *testing.T) {
	var err error = Arg("", nil)
	if err.Error() != "arg invalid" {
		t.Errorf(`err.Error() != "arg invalid" -- %q`, err.Error())
	}
	inner := Unwrap(err)
	if inner != nil {
		t.Error("inner != nil")
	}
}

func TestArgFields(t *testing.T) {
	var err error = ArgFields("foo", Ent("name", DataErrEmpty), Ent("bar", DataErrUnspecified))
	if err.Error() != "arg foo invalid: name: empty, bar: unspecified" {
		t.Errorf(`err.Error() != "arg foo invalid: name: empty, bar: unspecified" -- %q`, err.Error())
	}
}
