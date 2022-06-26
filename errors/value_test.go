package errors

import "testing"

func TestValueMalformedBare(t *testing.T) {
	var err error = ErrValueMalformed
	if err.Error() != "malformed" {
		t.Errorf(`err.Error() != "malformed" -- %q`, err.Error())
	}
	if err != ErrValueMalformed {
		t.Error("err != DataErrMalformed")
	}
	if !IsValueMalformedError(err) {
		t.Errorf("!IsValueMalformed(err)")
	}
}

func TestValueMalformedNil(t *testing.T) {
	var err error
	if IsValueMalformedError(err) {
		t.Error("IsValueMalformed(err)")
	}
}

func TestValueMalformedNegative(t *testing.T) {
	var err error = Msg("nothing")
	if IsValueMalformedError(err) {
		t.Error("IsValueMalformed(err)")
	}
}

func TestValueMalformedComplexNil(t *testing.T) {
	var err error = ValueMalformed(nil)
	if err.Error() != "malformed" {
		t.Errorf(`err.Error() != "malformed" -- %q`, err.Error())
	}
	if !IsValueMalformedError(err) {
		t.Error("!IsValueMalformed(err)")
	}
	if Unwrap(err) != nil {
		t.Errorf("Unwrap(err) == nil -- %#v", Unwrap(err))
	}
}

func TestValueMalformedComplexWithDetails(t *testing.T) {
	var err error = ValueMalformed(ErrValueInvalid)
	if err.Error() != "malformed: invalid" {
		t.Errorf(`err.Error() != "malformed: invalid" -- %q`, err.Error())
	}
	if !IsValueMalformedError(err) {
		t.Error("!IsValueMalformed(err)")
	}
	if Unwrap(err) == nil {
		t.Errorf("Unwrap(err) == nil -- %#v", Unwrap(err))
	}
}

func TestDescriptorDetailsBlank(t *testing.T) {
	var err error = &valueDescriptorDetailsError{}
	if err.Error() != "" {
		t.Errorf(`err.Error() != "" -- %q`, err.Error())
	}
}

func TestDescriptorDetailsOnlyDetails(t *testing.T) {
	var err error = &valueDescriptorDetailsError{details: ErrValueInvalid}
	if err.Error() != "invalid" {
		t.Errorf(`err.Error() != "invalid" -- %q`, err.Error())
	}
}
