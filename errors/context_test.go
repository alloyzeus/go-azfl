package errors

import "testing"

func TestContextEmpty(t *testing.T) {
	var err error = Context(nil)
	if err.Error() != "context invalid" {
		t.Errorf(`err.Error() != "context invalid" -- %q`, err.Error())
	}
	inner := Unwrap(err)
	if inner != nil {
		t.Error("inner != nil")
	}
}

func TestContextConstantDescriptor(t *testing.T) {
	var err error = Context(ErrValueUnspecified, nil)
	if err.Error() != "context unspecified" {
		t.Errorf(`err.Error() != "context unspecified" -- %q`, err.Error())
	}
	if d, ok := err.(hasDescriptor); !ok {
		t.Error("err.(hasDescriptor)")
	} else {
		desc := d.Descriptor()
		if desc == nil {
			t.Error("desc == nil")
		}
		if desc != ErrValueUnspecified {
			t.Error("desc != ErrValueUnspecified")
		}
	}
	if !IsContextUnspecifiedError(err) {
		t.Error("!IsContextUnspecified(err)")
	}
}

func TestContextWrappedConstantDescriptor(t *testing.T) {
	var err error = Context(DescWrap(ErrValueUnspecified, nil))
	if err.Error() != "context unspecified" {
		t.Errorf(`err.Error() != "context unspecified" -- %q`, err.Error())
	}
	if d, ok := err.(hasDescriptor); !ok {
		t.Error("err.(hasDescriptor)")
	} else {
		desc := d.Descriptor()
		if desc == nil {
			t.Error("desc == nil")
		}
		if desc != ErrValueUnspecified {
			t.Error("desc != ErrValueUnspecified")
		}
	}
	if !IsContextUnspecifiedError(err) {
		t.Error("!IsContextUnspecified(err)")
	}
}

func TestContextUnspecified(t *testing.T) {
	var err error = ContextUnspecified()
	if err.Error() != "context unspecified" {
		t.Errorf(`err.Error() != "context unspecified" -- %q`, err.Error())
	}
	if ctxErr, ok := err.(ContextError); !ok {
		t.Error("err.(ContextError)")
	} else {
		if ctxErr == nil {
			t.Error("argErr == nil")
		}
	}
	inner := Unwrap(err)
	if inner == nil {
		t.Error("inner == nil")
	}
	if inner != ErrValueUnspecified {
		t.Error("inner != DataErrUnspecified")
	}
	if d, ok := err.(hasDescriptor); !ok {
		t.Error("err.(hasDescriptor)")
	} else {
		desc := d.Descriptor()
		if desc == nil {
			t.Error("desc == nil")
		}
		if desc != ErrValueUnspecified {
			t.Error("desc != ErrValueUnspecified")
		}
	}
}

func TestContextUnspecifiedCustomBare(t *testing.T) {
	var err error = Context(ErrValueUnspecified)
	if !IsContextUnspecifiedError(err) {
		t.Errorf("!IsContextUnspecified(err)")
	}
}

// Ensure that the descriptor is for the context, not for others.
func TestContextUnspecifiedCustomWrap(t *testing.T) {
	var err error = Context(Wrap("", ErrValueUnspecified))
	if IsContextUnspecifiedError(err) {
		t.Errorf("IsContextUnspecified(err)")
	}
}

func TestContextFelds(t *testing.T) {
	var err error = ContextFields(Ent("authorization", ErrValueUnspecified))
	if err.Error() != "context invalid: authorization: unspecified" {
		t.Errorf(`err.Error() != "context invalid: authorization: unspecified" -- %q`, err.Error())
	}
}

func TestIsContextErrorNil(t *testing.T) {
	var err error
	if IsContextError(err) {
		t.Error("IsContextError(err)")
	}
}

func TestIsContextErrorNegative(t *testing.T) {
	var err error = ErrValueUnspecified
	if IsContextError(err) {
		t.Error("IsContextError(err)")
	}
}

func TestIsContextUnspecifiedErrorNil(t *testing.T) {
	var err error
	if IsContextUnspecifiedError(err) {
		t.Error("IsContextUnspecifiedError(err)")
	}
}

func TestIsContextUnspecifiedErrorNegative(t *testing.T) {
	var err error = ErrValueUnspecified
	if IsContextUnspecifiedError(err) {
		t.Error("IsContextUnspecifiedError(err)")
	}
}
