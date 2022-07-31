package errors

import "testing"

func TestAccessErrorBlank(t *testing.T) {
	var err error = Access()
	assert(t, "access error", err.Error())
	assert(t, true, IsCallError(err))
	assert(t, true, IsContextError(err))
	assert(t, true, IsAccessError(err))
	assert(t, false, IsAccessForbidden(err))
	assert(t, false, IsAuthorizationInvalid(err))
	assert(t, nil, Unwrap(err))
}

func TestAccessErrorNoDesc(t *testing.T) {
	var err error = Access().Wrap(N("foo").Desc(ErrEntityNotFound))
	assert(t, "access: foo: not found", err.Error())
}

func TestAccessDeniedEmpty(t *testing.T) {
	var err error = AccessForbidden()
	assert(t, "access forbidden", err.Error())
	assert(t, true, IsCallError(err))
	assert(t, true, IsContextError(err))
	assert(t, true, IsAccessError(err))
	assert(t, true, IsAccessForbidden(err))
	assert(t, false, IsAuthorizationInvalid(err))
	assert(t, nil, Unwrap(err))
}

func TestAccessDeniedEntity(t *testing.T) {
	var err error = AccessForbidden().Wrap(N("foo"))
	assert(t, "access forbidden: foo", err.Error())
	assert(t, true, IsAccessError(err))
	assert(t, true, IsAccessForbidden(err))
	assertNotEqual(t, nil, UnwrapDescriptor(err))
	assert(t, ErrAccessForbidden, UnwrapDescriptor(err))
}

func TestIsAccessDenied(t *testing.T) {
	var err error = ErrAccessForbidden
	assert(t, "forbidden", err.Error())
	assert(t, true, IsAccessForbidden(err))
	assert(t, true, IsAccessError(err))
}
