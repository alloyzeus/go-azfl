package errors

import "testing"

func TestArgUnspecifiedEmpty(t *testing.T) {
	var err error = ArgUnspecified("")
	assert(t, "arg unspecified", err.Error())

	argErr, ok := err.(ArgumentError)
	assert(t, true, ok)
	assertNotEqual(t, nil, argErr)
	assert(t, argErr, argErr.CallError())
	assert(t, "", argErr.ArgumentName())
	assert(t, true, IsArgumentUnspecifiedError(err))
	assert(t, ErrValueUnspecified, UnwrapDescriptor(err))

	wrapped := Unwrap(err)
	assert(t, nil, wrapped)

	desc := UnwrapDescriptor(err)
	assertNotEqual(t, nil, desc)
	assert(t, ErrValueUnspecified, desc)
}

func TestArgUnspecifiedFoo(t *testing.T) {
	var err error = ArgUnspecified("foo")
	assert(t, "arg foo: unspecified", err.Error())
	assert(t, true, IsArgumentUnspecifiedError(err))
	assert(t, true, IsArgumentUnspecified(err, "foo"))
	assert(t, false, IsArgumentUnspecified(err, "bar"))

	argErr, ok := err.(ArgumentError)
	assert(t, true, ok)
	assertNotEqual(t, nil, argErr)
	assert(t, "foo", argErr.ArgumentName())
	assert(t, true, IsArgumentUnspecifiedError(err))
	assert(t, ErrValueUnspecified, UnwrapDescriptor(err))

	wrapped := Unwrap(err)
	assert(t, nil, wrapped)

	desc := UnwrapDescriptor(err)
	assertNotEqual(t, nil, desc)
	assert(t, ErrValueUnspecified, desc)
}

func TestIsArgUnspecifiedErrorNil(t *testing.T) {
	var err error
	assert(t, false, IsArgumentUnspecifiedError(err))
	assert(t, false, IsArgumentUnspecified(err, "foo"))
}

func TestIsArgUnspecifiedNegative(t *testing.T) {
	var err error = ErrValueInvalid
	assert(t, false, IsArgumentUnspecified(err, "foo"))
}

func TestIsArgUnspecifiedWrongArgName(t *testing.T) {
	var err error = ArgUnspecified("foo")
	assert(t, false, IsArgumentUnspecified(err, "bar"))
}

func TestIsArgUnspecifiedCustomStruct(t *testing.T) {
	var err error = &customArgError{argName: "foo"}
	assert(t, false, IsArgumentUnspecified(err, "foo"))
}

func TestArgEmpty(t *testing.T) {
	var err error = Arg("")
	assert(t, "arg error", err.Error())

	inner := Unwrap(err)
	assert(t, nil, inner)
	assert(t, false, IsArgumentUnspecifiedError(err))
}

func TestArgFields(t *testing.T) {
	var err error = Arg("foo").Fieldset(
		Ent("name").Desc(ErrValueEmpty), Ent("bar").Desc(ErrValueUnspecified))
	assert(t, "arg foo: name: empty, bar: unspecified", err.Error())
}

func TestArgFieldsNoName(t *testing.T) {
	var err error = Arg("").Fieldset(
		Ent("name").Desc(ErrValueEmpty), Ent("bar").Desc(ErrValueUnspecified))
	assert(t, "arg: name: empty, bar: unspecified", err.Error())
}

func TestArg1(t *testing.T) {
	var err error = Arg1().Desc(ErrValueUnspecified)
	assert(t, "arg unspecified", err.Error())
}

func TestArgDescWrap(t *testing.T) {
	inner := Msg("inner")
	var err error = Arg("foo").Desc(ErrValueMalformed).Wrap(inner)
	assert(t, "arg foo: malformed: inner", err.Error())

	wrapped := Unwrap(err)
	assertNotEqual(t, nil, wrapped)
	assert(t, inner, wrapped)

	assert(t, true, HasDescriptor(err, ErrValueMalformed))

	desc := UnwrapDescriptor(err)
	assertNotEqual(t, nil, desc)
	assert(t, ErrValueMalformed, desc)
}

func TestArgDescMsg(t *testing.T) {
	const customDesc = "custom descriptor"
	var err error = Arg("saturn").DescMsg(customDesc)
	assert(t, "arg saturn: custom descriptor", err.Error())
	assert(t, true, HasDescriptorText(err, customDesc))

	desc := UnwrapDescriptor(err)
	assertNotEqual(t, nil, desc)
	assert(t, customDesc, desc.Error())
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
