package errors

import "testing"

func TestDataMalformedBare(t *testing.T) {
	var err error = ErrDataMalformed
	if err != ErrDataMalformed {
		t.Error("err != DataErrMalformed")
	}
	if !IsDataMalformedError(err) {
		t.Errorf("!IsDataMalformed(err)")
	}
	if err.Error() != "malformed" {
		t.Errorf(`err.Error() != "malformed" -- %q`, err.Error())
	}
}

func TestDataMalformedNil(t *testing.T) {
	var err error
	if IsDataMalformedError(err) {
		t.Error("IsDataMalformed(err)")
	}
}

func TestDataMalformedNegative(t *testing.T) {
	var err error = Msg("nothing")
	if IsDataMalformedError(err) {
		t.Error("IsDataMalformed(err)")
	}
}

func TestDataMalformedComplexNil(t *testing.T) {
	var err error = DataMalformed(nil)
	if !IsDataMalformedError(err) {
		t.Error("!IsDataMalformed(err)")
	}
	if err.Error() != "malformed" {
		t.Errorf(`err.Error() != "malformed" -- %q`, err.Error())
	}
	if Unwrap(err) != nil {
		t.Errorf("Unwrap(err) == nil -- %#v", Unwrap(err))
	}
}

func TestDataMalformedComplexWithDetails(t *testing.T) {
	var err error = DataMalformed(ErrDataInvalid)
	if !IsDataMalformedError(err) {
		t.Error("!IsDataMalformed(err)")
	}
	if err.Error() != "malformed: invalid" {
		t.Errorf(`err.Error() != "malformed: invalid" -- %q`, err.Error())
	}
	if Unwrap(err) == nil {
		t.Errorf("Unwrap(err) == nil -- %#v", Unwrap(err))
	}
}

func TestDescriptorDetailsBlank(t *testing.T) {
	var err error = &descriptorDetailsError{}
	if err.Error() != "" {
		t.Errorf(`err.Error() != "" -- %q`, err.Error())
	}
}

func TestDescriptorDetailsOnlyDetails(t *testing.T) {
	var err error = &descriptorDetailsError{details: ErrDataInvalid}
	if err.Error() != "invalid" {
		t.Errorf(`err.Error() != "invalid" -- %q`, err.Error())
	}
}
