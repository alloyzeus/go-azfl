package errors

import "testing"

func TestUnwrapErrorSetNil(t *testing.T) {
	var err error
	errs := UnwrapErrorSet(err)
	//assert(t, nil, errs)
	assert(t, 0, len(errs))
}

func TestUnwrapErrorSetWrongType(t *testing.T) {
	var err error = ErrValueMalformed
	errs := UnwrapErrorSet(err)
	//assert(t, nil, errs)
	assert(t, 0, len(errs))
}

func TestErrorSetEmpty(t *testing.T) {
	var err error = Set()
	assert(t, "", err.Error())

	errs := UnwrapErrorSet(err)
	assert(t, 0, len(errs))
}

func TestErrorSetTwo(t *testing.T) {
	var err error = Set(Msg("first"), Msg("second"))
	assert(t, "first, second", err.Error())

	errs := UnwrapErrorSet(err)
	assert(t, 2, len(errs))

	errSet, ok := err.(ErrorSet)
	assert(t, true, ok)
	assert(t, 2, len(errSet.Errors()))
}
