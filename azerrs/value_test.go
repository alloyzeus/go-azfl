package errors

import "testing"

func TestValueMalformedBare(t *testing.T) {
	var err error = ErrValueMalformed
	assert(t, "malformed", err.Error())
	assert(t, ErrValueMalformed, err)
	assert(t, true, IsValueMalformedError(err))
}

func TestValueMalformedNil(t *testing.T) {
	var err error
	assert(t, false, IsValueMalformedError(err))
}

func TestValueMalformedNegative(t *testing.T) {
	var err error = Msg("nothing")
	assert(t, false, IsValueMalformedError(err))
}

func TestValueMalformedComplexNil(t *testing.T) {
	var err error = ValueMalformed()
	assert(t, "malformed", err.Error())
	assert(t, true, IsValueMalformedError(err))
	assert(t, nil, Unwrap(err))
}

func TestValueMalformedComplexWithDetails(t *testing.T) {
	var err error = ValueMalformed().Wrap(ErrValueInvalid)
	assert(t, "malformed: invalid", err.Error())
	assert(t, true, IsValueMalformedError(err))
	assert(t, ErrValueInvalid, Unwrap(err))
}

func TestDescriptorDetailsBlank(t *testing.T) {
	var err error = &valueDescriptorDetailsError{}
	assert(t, "", err.Error())
}

func TestDescriptorDetailsOnlyDetails(t *testing.T) {
	var err error = &valueDescriptorDetailsError{wrapped: ErrValueInvalid}
	assert(t, "invalid", err.Error())
}
