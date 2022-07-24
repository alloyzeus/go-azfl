package errors

import "testing"

func TestContextEmpty(t *testing.T) {
	var err error = Context()
	assert(t, "context invalid", err.Error())

	inner := Unwrap(err)
	assert(t, nil, inner)

	desc := UnwrapDescriptor(err)
	assert(t, nil, desc)
}

func TestContextConstantDescriptor(t *testing.T) {
	var err error = Context().Desc(ErrValueUnspecified)
	assert(t, "context unspecified", err.Error())

	d, ok := err.(hasDescriptor)
	assert(t, true, ok)

	desc := d.Descriptor()
	assertNotEqual(t, nil, desc)
	assert(t, ErrValueUnspecified, desc)

	assert(t, true, IsContextUnspecifiedError(err))
}

func TestContextDescriptor(t *testing.T) {
	var err error = Context().Desc(ErrValueUnspecified)
	assert(t, "context unspecified", err.Error())
	if d, ok := err.(hasDescriptor); !ok {
		t.Error("err.(hasDescriptor)")
	} else {
		desc := d.Descriptor()
		if desc == nil {
			t.Error("desc == nil")
		}
		if desc != ErrValueUnspecified {
			t.Error("desc != ErrValueUnspecified")
		}
	}
	if !IsContextUnspecifiedError(err) {
		t.Error("!IsContextUnspecified(err)")
	}
}

func TestContextDescWrap(t *testing.T) {
	var err error = Context().Desc(ErrValueMalformed).Wrap(Msg("missing authorization"))
	assert(t, "context malformed: missing authorization", err.Error())
}

func TestContextUnspecified(t *testing.T) {
	var err error = ContextUnspecified()
	assert(t, "context unspecified", err.Error())

	ctxErr, ok := err.(ContextError)
	assert(t, true, ok)
	assert(t, ctxErr, ctxErr.CallError())
	assert(t, ctxErr, ctxErr.ContextError())
	assertNotEqual(t, nil, ctxErr)

	inner := Unwrap(err)
	assertNotEqual(t, nil, inner)
	assert(t, ErrValueUnspecified, inner)

	d, ok := err.(hasDescriptor)
	assert(t, true, ok)

	desc := d.Descriptor()
	assertNotEqual(t, nil, desc)
	assert(t, ErrValueUnspecified, desc)
}

func TestContextUnspecifiedCustomBare(t *testing.T) {
	var err error = Context().Wrap(ErrValueUnspecified)
	assert(t, true, IsContextUnspecifiedError(err))
}

func TestContextFelds(t *testing.T) {
	var err error = Context().Fieldset(N("authorization").Desc(ErrValueUnspecified))
	assert(t, "context invalid: authorization: unspecified", err.Error())
}

func TestIsContextErrorNil(t *testing.T) {
	var err error
	assert(t, false, IsContextError(err))
}

func TestIsContextErrorNegative(t *testing.T) {
	var err error = ErrValueUnspecified
	assert(t, false, IsContextError(err))
}

func TestIsContextUnspecifiedErrorNil(t *testing.T) {
	var err error
	assert(t, false, IsContextUnspecifiedError(err))
}

func TestIsContextUnspecifiedErrorNegative(t *testing.T) {
	var err error = ErrValueUnspecified
	assert(t, false, IsContextUnspecifiedError(err))
}
