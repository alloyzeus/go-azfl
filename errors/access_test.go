package errors

import "testing"

func TestAccessErrorBlank(t *testing.T) {
	var err error = Access(nil, nil)
	if err.Error() != "access error" {
		t.Errorf(`err.Error() != "access error" -- %q`, err.Error())
	}
	if !IsCallError(err) {
		t.Error("!IsCallError(err)")
	}
	if !IsContextError(err) {
		t.Error("!IsContextError(err)")
	}
	if !IsAccessError(err) {
		t.Error("!IsAccessError(err)")
	}
	if IsAccessDenied(err) {
		t.Error("!IsAccessDenied(err)")
	}
	if IsAuthorizationInvalid(err) {
		t.Error("IsAuthorizationInvalid(err)")
	}
	if Unwrap(err) != nil {
		t.Error("Unwrap(err) != nil")
	}
}

func TestAccessErrorNoDesc(t *testing.T) {
	var err error = Access(nil, EntNotFound("foo", nil))
	if err.Error() != "access: foo: not found" {
		t.Errorf(`err.Error() != "access: foo: not found" -- %q`, err.Error())
	}
}

func TestAccessDeniedEmpty(t *testing.T) {
	var err error = AccessDenied(nil)
	if err.Error() != "access denied" {
		t.Errorf(`err.Error() != "access denied" -- %q`, err.Error())
	}
	if !IsCallError(err) {
		t.Error("!IsCallError(err)")
	}
	if !IsContextError(err) {
		t.Error("!IsContextError(err)")
	}
	if !IsAccessError(err) {
		t.Error("!IsAccessError(err)")
	}
	if !IsAccessDenied(err) {
		t.Error("!IsAccessDenied(err)")
	}
	if IsAuthorizationInvalid(err) {
		t.Error("IsAuthorizationInvalid(err)")
	}
	if Unwrap(err) != nil {
		t.Error("Unwrap(err) != nil")
	}
}

func TestAccessDeniedEntity(t *testing.T) {
	var err error = AccessDenied(Ent("foo", nil))
	if err.Error() != "access denied: foo" {
		t.Errorf(`err.Error() != "access denied: foo" -- %q`, err.Error())
	}
	if !IsAccessDenied(err) {
		t.Error("!IsAccessDenied(err)")
	}
}

func TestIsAccessDenied(t *testing.T) {
	var err error = ErrAccessDenied
	if err.Error() != "denied" {
		t.Errorf(`err.Error() != "denied" -- %q`, err.Error())
	}
	if !IsAccessDenied(err) {
		t.Error("!IsAccessDenied(err)")
	}
}
