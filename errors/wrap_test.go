package errors

import "testing"

func TestWrapEmpty(t *testing.T) {
	var err error = Wrap("", nil)
	if err.Error() != "" {
		t.Errorf(`err.Error() != "" -- %q`, err.Error())
	}

	if u, ok := err.(Unwrappable); !ok {
		t.Errorf("err.(Unwrappable)")
	} else {
		inner := u.Unwrap()
		if inner != nil {
			t.Errorf("inner != nil")
		}
		if inner == ErrValueUnspecified {
			t.Errorf("inner == DataErrUnspecified")
		}
	}

	inner := Unwrap(err)
	if inner != nil {
		t.Errorf("inner != nil")
	}
	if inner == ErrValueUnspecified {
		t.Errorf("inner == DataErrUnspecified")
	}
}

func TestWrapNoContext(t *testing.T) {
	var err error = Wrap("", ErrValueUnspecified)
	if err.Error() != "unspecified" {
		t.Errorf(`err.Error() != "unspecified" -- %q`, err.Error())
	}

	if u, ok := err.(Unwrappable); !ok {
		t.Errorf("err.(Unwrappable)")
	} else {
		inner := u.Unwrap()
		if inner == nil {
			t.Errorf("inner == nil")
		}
		if inner != ErrValueUnspecified {
			t.Errorf("inner != DataErrUnspecified")
		}
	}

	inner := Unwrap(err)
	if inner == nil {
		t.Errorf("inner == nil")
	}
	if inner != ErrValueUnspecified {
		t.Errorf("inner != DataErrUnspecified")
	}
}

func TestWrapNoErr(t *testing.T) {
	var err error = Wrap("unknown", nil)
	if err.Error() != "unknown" {
		t.Errorf(`err.Error() != "unknown" -- %q`, err.Error())
	}

	if u, ok := err.(Unwrappable); !ok {
		t.Errorf("err.(Unwrappable)")
	} else {
		inner := u.Unwrap()
		if inner != nil {
			t.Errorf("inner != nil")
		}
		if inner == ErrValueUnspecified {
			t.Errorf("inner == DataErrUnspecified")
		}
	}

	inner := Unwrap(err)
	if inner != nil {
		t.Errorf("inner != nil")
	}
	if inner == ErrValueUnspecified {
		t.Errorf("inner == DataErrUnspecified")
	}
}

func TestWrap(t *testing.T) {
	var err error = Wrap("launch code", ErrValueUnspecified)
	if err.Error() != "launch code: unspecified" {
		t.Errorf(`err.Error() != "launch code: unspecified" -- %q`, err.Error())
	}

	if u, ok := err.(Unwrappable); !ok {
		t.Errorf("err.(Unwrappable)")
	} else {
		inner := u.Unwrap()
		if inner == nil {
			t.Errorf("inner == nil")
		}
		if inner != ErrValueUnspecified {
			t.Errorf("inner != DataErrUnspecified")
		}
	}

	inner := Unwrap(err)
	if inner == nil {
		t.Errorf("inner == nil")
	}
	if inner != ErrValueUnspecified {
		t.Errorf("inner != DataErrUnspecified")
	}
}
