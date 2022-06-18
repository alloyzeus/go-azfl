package errors

import "testing"

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
	var err error = Ent("", ErrDataMalformed)
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
		Ent("name", ErrDataEmpty),
		Ent("age", ErrDataUnspecified))
	if err.Error() != "user invalid: name: empty, age: unspecified" {
		t.Errorf(`err.Error() != "user invalid: name: empty, age: unspecified" -- %q`, err.Error())
	}
}

func TestEntNotFound(t *testing.T) {
	var fooNotFound error = EntNotFound("foo", nil)
	if !IsEntNotFoundError(fooNotFound) {
		t.Errorf("!IsEntNotFound(fooNotFound)")
	}
	if fooNotFound.Error() != "foo: not found" {
		t.Errorf(`fooNotFound.Error() != "foo: not found" -- %q`, fooNotFound.Error())
	}
	var notFoundBare error = ErrEntityNotFound
	if IsEntNotFoundError(notFoundBare) {
		t.Errorf("IsEntNotFound(notFoundBare)")
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
	if !IsEntInvalidError(fooInvalid) {
		t.Errorf("!IsEntInvalid(fooInvalid)")
	}
	if fooInvalid.Error() != "foo: invalid" {
		t.Errorf(`fooInvalid.Error() != "foo: invalid" -- %q`, fooInvalid.Error())
	}
	var notFoundBare error = ErrDataInvalid
	if IsEntInvalidError(notFoundBare) {
		t.Errorf("IsEntInvalid(notFoundBare)")
	}
	if IsArgumentError(fooInvalid) {
		t.Errorf("IsArgumentError(fooInvalid)")
	}
	if IsCallError(fooInvalid) {
		t.Errorf("IsCallError(fooInvalid)")
	}
}
