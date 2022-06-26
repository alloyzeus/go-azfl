package errors

import "testing"

func TestEmpty(t *testing.T) {
	var err error = DescWrap(nil, nil)
	if err.Error() != "" {
		t.Errorf(`err.Error() != "" -- %q`, err.Error())
	}
}

func TestDescriptorOnly(t *testing.T) {
	var err error = DescWrap(ErrValueEmpty, nil)
	if err.Error() != "empty" {
		t.Errorf(`err.Error() != "empty" -- %q`, err.Error())
	}
	if UnwrapDescriptor(err) != ErrValueEmpty {
		t.Error("UnwrapDescriptor(err) != ErrValueEmpty")
	}
	if Unwrap(err) != nil {
		t.Error("Unwrap(err) != nil")
	}
}

func TestDetailsOnly(t *testing.T) {
	var err error = DescWrap(nil, Msg("unexpected condition"))
	if err.Error() != "unexpected condition" {
		t.Errorf(`err.Error() != "unexpected condition" -- %q`, err.Error())
	}
}

func TestDescWrapSimple(t *testing.T) {
	var err error = DescWrap(ErrAccessDenied, Msg("insufficient permission"))
	if err.Error() != "denied: insufficient permission" {
		t.Errorf(`err.Error() != "denied: insufficient permission" -- %q`, err.Error())
	}
}
