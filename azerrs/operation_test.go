package errors

import "testing"

func TestOpEmpty(t *testing.T) {
	var err error = Op("")

	assert(t, "operation error", err.Error())
	assert(t, nil, Unwrap(err))

	opErr, ok := err.(OperationError)
	assert(t, true, ok)
	assertNotEqual(t, nil, opErr)
	assert(t, "", opErr.OperationName())
}

func TestOpJustName(t *testing.T) {
	var err error = Op("DeleteAccount")

	assert(t, "DeleteAccount error", err.Error())
	assert(t, nil, Unwrap(err))

	opErr, ok := err.(OperationError)
	assert(t, true, ok)
	assertNotEqual(t, nil, opErr)
	assert(t, "DeleteAccount", opErr.OperationName())
}

func TestOpJustWrapped(t *testing.T) {
	var err error = Op("").Wrap(ErrAccessForbidden)

	assert(t, "operation error: forbidden", err.Error())

	inner := Unwrap(err)
	assertNotEqual(t, nil, inner)
	assert(t, ErrAccessForbidden, inner)

	opErr, ok := err.(OperationError)
	assert(t, true, ok)
	assertNotEqual(t, nil, opErr)
	assert(t, "", opErr.OperationName())
	assert(t, ErrAccessForbidden, opErr.Unwrap())
}

func TestOpNameAndWrapped(t *testing.T) {
	var err error = Op("DropDatabase").Wrap(ErrAccessForbidden)

	assert(t, "DropDatabase: forbidden", err.Error())

	inner := Unwrap(err)
	assertNotEqual(t, nil, inner)
	assert(t, ErrAccessForbidden, inner)

	opErr, ok := err.(OperationError)
	assert(t, true, ok)
	assertNotEqual(t, nil, opErr)
	assert(t, "DropDatabase", opErr.OperationName())
	assert(t, ErrAccessForbidden, opErr.Unwrap())
}
