package errors

import "testing"

func TestWrapEmpty(t *testing.T) {
	var err error = Wrap("", nil)
	if u, ok := err.(Unwrappable); !ok {
		t.Errorf("err.(Unwrappable)")
	} else {
		inner := u.Unwrap()
		if inner != nil {
			t.Errorf("inner != nil")
		}
		if inner == ErrDataUnspecified {
			t.Errorf("inner == DataErrUnspecified")
		}
	}

	inner := Unwrap(err)
	if inner != nil {
		t.Errorf("inner != nil")
	}
	if inner == ErrDataUnspecified {
		t.Errorf("inner == DataErrUnspecified")
	}

	if err.Error() != "" {
		t.Errorf(`err.Error() != "" -- %q`, err.Error())
	}
}

func TestWrapNoContext(t *testing.T) {
	var err error = Wrap("", ErrDataUnspecified)
	if u, ok := err.(Unwrappable); !ok {
		t.Errorf("err.(Unwrappable)")
	} else {
		inner := u.Unwrap()
		if inner == nil {
			t.Errorf("inner == nil")
		}
		if inner != ErrDataUnspecified {
			t.Errorf("inner != DataErrUnspecified")
		}
	}

	inner := Unwrap(err)
	if inner == nil {
		t.Errorf("inner == nil")
	}
	if inner != ErrDataUnspecified {
		t.Errorf("inner != DataErrUnspecified")
	}

	if err.Error() != "unspecified" {
		t.Errorf(`err.Error() != "unspecified" -- %q`, err.Error())
	}
}

func TestWrapNoErr(t *testing.T) {
	var err error = Wrap("unknown", nil)
	if u, ok := err.(Unwrappable); !ok {
		t.Errorf("err.(Unwrappable)")
	} else {
		inner := u.Unwrap()
		if inner != nil {
			t.Errorf("inner != nil")
		}
		if inner == ErrDataUnspecified {
			t.Errorf("inner == DataErrUnspecified")
		}
	}

	inner := Unwrap(err)
	if inner != nil {
		t.Errorf("inner != nil")
	}
	if inner == ErrDataUnspecified {
		t.Errorf("inner == DataErrUnspecified")
	}

	if err.Error() != "unknown" {
		t.Errorf(`err.Error() != "unknown" -- %q`, err.Error())
	}
}

func TestWrap(t *testing.T) {
	var err error = Wrap("launch code", ErrDataUnspecified)
	if u, ok := err.(Unwrappable); !ok {
		t.Errorf("err.(Unwrappable)")
	} else {
		inner := u.Unwrap()
		if inner == nil {
			t.Errorf("inner == nil")
		}
		if inner != ErrDataUnspecified {
			t.Errorf("inner != DataErrUnspecified")
		}
	}

	inner := Unwrap(err)
	if inner == nil {
		t.Errorf("inner == nil")
	}
	if inner != ErrDataUnspecified {
		t.Errorf("inner != DataErrUnspecified")
	}

	if err.Error() != "launch code: unspecified" {
		t.Errorf(`err.Error() != "launch code: unspecified" -- %q`, err.Error())
	}
}
