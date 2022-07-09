package errors

import "testing"

func TestArgUnspecifiedEmpty(t *testing.T) {
	var err error = ArgUnspecified("")
	assert(t, "arg unspecified", err.Error())

	argErr, ok := err.(ArgumentError)
	assert(t, true, ok)
	assertNotEqual(t, nil, argErr)
	assert(t, "", argErr.ArgumentName())
	assert(t, true, IsArgUnspecifiedError(err))
	assert(t, ErrValueUnspecified, UnwrapDescriptor(err))
}

func TestArgUnspecifiedFoo(t *testing.T) {
	var err error = ArgUnspecified("foo")
	assert(t, "arg foo: unspecified", err.Error())

	argErr, ok := err.(ArgumentError)
	assert(t, true, ok)
	assertNotEqual(t, nil, argErr)
	assert(t, "foo", argErr.ArgumentName())
	assert(t, true, IsArgUnspecifiedError(err))
	assert(t, ErrValueUnspecified, UnwrapDescriptor(err))

	inner := Unwrap(err)
	assertNotEqual(t, nil, inner)
	assert(t, ErrValueUnspecified, inner)
	assert(t, true, IsArgUnspecifiedError(err))
	assert(t, true, IsArgUnspecified(err, "foo"))
	assert(t, false, IsArgUnspecified(err, "bar"))
}

func TestIsArgUnspecifiedErrorNil(t *testing.T) {
	var err error
	assert(t, false, IsArgUnspecifiedError(err))
	assert(t, false, IsArgUnspecified(err, "foo"))
}

func TestIsArgUnspecifiedNegative(t *testing.T) {
	var err error = ErrValueInvalid
	assert(t, false, IsArgUnspecified(err, "foo"))
}

func TestIsArgUnspecifiedWrongArgName(t *testing.T) {
	var err error = ArgUnspecified("foo")
	assert(t, false, IsArgUnspecified(err, "bar"))
}

func TestIsArgUnspecifiedCustomStruct(t *testing.T) {
	var err error = &customArgError{argName: "foo"}
	assert(t, false, IsArgUnspecified(err, "foo"))
}

func TestArgEmpty(t *testing.T) {
	var err error = Arg("", nil)
	assert(t, "arg error", err.Error())

	inner := Unwrap(err)
	assert(t, nil, inner)
	assert(t, false, IsArgUnspecifiedError(err))
}

func TestArgFields(t *testing.T) {
	var err error = ArgFields("foo", Ent("name", ErrValueEmpty), Ent("bar", ErrValueUnspecified))
	assert(t, "arg foo: name: empty, bar: unspecified", err.Error())
}

func TestArgFieldsNoName(t *testing.T) {
	var err error = ArgFields("", Ent("name", ErrValueEmpty), Ent("bar", ErrValueUnspecified))
	assert(t, "arg: name: empty, bar: unspecified", err.Error())
}

func TestArg1(t *testing.T) {
	var err error = Arg1(ErrValueUnspecified)
	assert(t, "arg unspecified", err.Error())
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
