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

	assert(t, "operation forbidden", err.Error())

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

func TestOpParamsBasic(t *testing.T) {
	var err error = Op("LaunchNuke").Params(N("authorizationCode").Desc(ErrValueMismatch))

	assert(t, "LaunchNuke error. authorizationCode: mismatch", err.Error())
}

func TestOpParamsRes(t *testing.T) {
	var err error = Op("Fetch").Params(N("part").Val("HEAD")).Wrap(ErrEntityUnreachable)

	//TODO: param values are quoted when necessary
	assert(t, "Fetch: unreachable. part=HEAD", err.Error())
}

func TestOpParamsNoWrap(t *testing.T) {
	var err error = Op("Fetch").Params(N("part").Val("HEAD"))

	assert(t, "Fetch error. part=HEAD", err.Error())
}

func TestOpNoNameParamsNoWrap(t *testing.T) {
	var err error = Op("").Params(N("part").Val("HEAD"))

	assert(t, "operation error. part=HEAD", err.Error())
}

func TestOpHintWrap(t *testing.T) {
	var err error = Op("SwimmingRevolution").Params(N("location").Val("Pacific Ocean")).
		Wrap(Msg("sun is shining")).Hint("Please check the water temperature first")

	assert(t, `SwimmingRevolution: sun is shining. location="Pacific Ocean". Please check the water temperature first`, err.Error())
}

func TestOpHintNoWrap(t *testing.T) {
	var err error = Op("SwimmingRevolution").Params(N("location").Val("Pacific Ocean")).
		Hint("Please check the water temperature first")

	assert(t, `SwimmingRevolution error. location="Pacific Ocean". Please check the water temperature first`, err.Error())
}
