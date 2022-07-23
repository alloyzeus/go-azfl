package errors

import "testing"

func TestEmpty(t *testing.T) {
	var err error = descWrap(nil, nil)
	assert(t, "", err.Error())
	assert(t, nil, Unwrap(err))
	assert(t, nil, UnwrapDescriptor(err))
}

func TestDescriptorOnly(t *testing.T) {
	var err error = descWrap(ErrValueEmpty, nil)
	assert(t, "empty", err.Error())
	assert(t, true, HasDescriptor(err, ErrValueEmpty))
	assert(t, false, HasDescriptor(err, ErrValueMalformed))
	assert(t, ErrValueEmpty, UnwrapDescriptor(err))
	assert(t, nil, Unwrap(err))
}

func TestDetailsOnly(t *testing.T) {
	var err error = descWrap(nil, Msg("unexpected condition"))
	assert(t, "unexpected condition", err.Error())
}

func TestDescWrapSimple(t *testing.T) {
	var err error = descWrap(ErrAccessForbidden, Msg("insufficient permission"))
	assert(t, "forbidden: insufficient permission", err.Error())
}
