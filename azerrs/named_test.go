package errors

import "testing"

//----

func TestNamedValueZero(t *testing.T) {
	var err error = N("")

	assert(t, "", err.Error())
	assert(t, nil, UnwrapDescriptor(err))
}

func TestNamedValueWrap(t *testing.T) {
	var err error = N("").Wrap(ErrValueEmpty)

	assert(t, "empty", err.Error())
}

func TestNamedValueWrapDesc(t *testing.T) {
	var err error = N("foo").Wrap(ErrValueEmpty)

	assert(t, "foo: empty", err.Error())
	assert(t, ErrValueEmpty, UnwrapDescriptor(err))
}

func TestNamedValueParam(t *testing.T) {
	var err error = N("").Fieldset(N("foo"))

	assert(t, "foo", err.Error())
}

func TestNamedValueHasSpace(t *testing.T) {
	var err error = N("yes").Val("has space")

	assert(t, `yes="has space"`, err.Error())

	namedError, ok := err.(NamedError)
	assert(t, true, ok)
	assertNotEqual(t, nil, namedError)
	assert(t, "has space", namedError.Value())

	rewrapped := N("no").Rewrap(err)
	assert(t, `no="has space"`, rewrapped.Error())
}

//----

func TestNamedValueUnsupportedNoName(t *testing.T) {
	var err error = NamedValueUnsupported("")

	//TODO: should be "value unsupported"
	assert(t, "unsupported", err.Error())
	assert(t, nil, Unwrap(err))
	assert(t, ErrValueUnsupported, UnwrapDescriptor(err))

	entErr, ok := err.(NamedError)
	assert(t, true, ok)
	assertNotEqual(t, nil, entErr)
	assert(t, "", entErr.Name())
}

func TestNamedValueUnsupportedFoo(t *testing.T) {
	var err error = NamedValueUnsupported("foo")
	assert(t, "foo: unsupported", err.Error())

	entErr, ok := err.(NamedError)
	assert(t, true, ok)
	assertNotEqual(t, nil, entErr)
	assert(t, "foo", entErr.Name())
	assert(t, ErrValueUnsupported, UnwrapDescriptor(err))

	wrapped := Unwrap(err)
	assert(t, nil, wrapped)

	desc := UnwrapDescriptor(err)
	assertNotEqual(t, nil, desc)
	assert(t, ErrValueUnsupported, desc)
}

//----

func TestNamedValueMalformedNoName(t *testing.T) {
	var err error = NamedValueMalformed("")

	//TODO: should be "entity value unsupported"
	assert(t, "malformed", err.Error())
	assert(t, nil, Unwrap(err))
	assert(t, ErrValueMalformed, UnwrapDescriptor(err))

	entErr, ok := err.(NamedError)
	assert(t, true, ok)
	assertNotEqual(t, nil, entErr)
	assert(t, "", entErr.Name())
}

func TestNamedValueMalformedFoo(t *testing.T) {
	var err error = NamedValueMalformed("foo")
	assert(t, "foo: malformed", err.Error())

	entErr, ok := err.(NamedError)
	assert(t, true, ok)
	assertNotEqual(t, nil, entErr)
	assert(t, "foo", entErr.Name())
	assert(t, ErrValueMalformed, UnwrapDescriptor(err))

	wrapped := Unwrap(err)
	assert(t, nil, wrapped)

	desc := UnwrapDescriptor(err)
	assertNotEqual(t, nil, desc)
	assert(t, ErrValueMalformed, desc)
}
