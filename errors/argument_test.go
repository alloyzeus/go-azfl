package errors

import "testing"

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
	if !IsArgUnspecifiedError(err) {
		t.Error("!IsArgUnspecifiedError(err)")
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
	if inner != ErrValueUnspecified {
		t.Error("inner != DataErrUnspecified")
	}
	if !IsArgUnspecifiedError(err) {
		t.Error("!IsArgUnspecified(err")
	}
	if !IsArgUnspecified(err, "foo") {
		t.Error(`!IsArgUnspecified(err, "foo")`)
	}
}

func TestIsArgUnspecifiedErrorNil(t *testing.T) {
	var err error
	if IsArgUnspecifiedError(err) {
		t.Error("IsArgUnspecifiedError(err)")
	}
}

func TestIsArgUnspecifiedNil(t *testing.T) {
	var err error
	if IsArgUnspecified(err, "foo") {
		t.Error(`IsArgUnspecified(err, "foo")`)
	}
}

func TestIsArgUnspecifiedNegative(t *testing.T) {
	var err error = ErrValueInvalid
	if IsArgUnspecified(err, "foo") {
		t.Error(`IsArgUnspecified(err, "foo")`)
	}
}

func TestIsArgUnspecifiedWrongArgName(t *testing.T) {
	var err error = ArgUnspecified("foo")
	if IsArgUnspecified(err, "bar") {
		t.Error(`IsArgUnspecified(err, "bar")`)
	}
}

func TestIsArgUnspecifiedCustomStruct(t *testing.T) {
	var err error = &customArgError{argName: "foo"}
	if IsArgUnspecified(err, "foo") {
		t.Error(`IsArgUnspecified(err, "foo")`)
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
	if IsArgUnspecifiedError(err) {
		t.Error("IsArgUnspecifiedError(err)")
	}
}

func TestArgFields(t *testing.T) {
	var err error = ArgFields("foo", Ent("name", ErrValueEmpty), Ent("bar", ErrValueUnspecified))
	if err.Error() != "arg foo invalid: name: empty, bar: unspecified" {
		t.Errorf(`err.Error() != "arg foo invalid: name: empty, bar: unspecified" -- %q`, err.Error())
	}
}

type customArgError struct {
	argName string
}

var (
	_ ArgumentError = &customArgError{}
)

func (e *customArgError) ArgumentName() string { return e.argName }
func (e *customArgError) Error() string        { return "custom arg error" }
func (e *customArgError) CallError() CallError { return e }
func (e *customArgError) Unwrap() error        { return nil }
