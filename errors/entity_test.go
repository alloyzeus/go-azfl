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
	var err error = Ent("", ErrValueMalformed)
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
		Ent("name", ErrValueEmpty),
		Ent("age", ErrValueUnspecified))
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

func TestIsEntNotFoundErrorCustomNegative(t *testing.T) {
	var err error = &customEntError{entID: "foo"}
	if IsEntNotFoundError(err) {
		t.Error(`IsEntNotFoundError(err)`)
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
	var notFoundBare error = ErrValueInvalid
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

func TestIsEntInvalidErrorCustomNegative(t *testing.T) {
	var err error = &customEntError{entID: "foo"}
	if IsEntInvalidError(err) {
		t.Error(`IsEntInvalidError(err)`)
	}
}

type customEntError struct {
	entID string
}

var (
	_ EntityError = &customEntError{}
)

func (e *customEntError) EntityIdentifier() string { return e.entID }
func (e *customEntError) Error() string            { return "custom ent error" }
func (e *customEntError) CallError() CallError     { return e }
func (e *customEntError) Unwrap() error            { return nil }
