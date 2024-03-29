package errors

import "testing"

func TestEntityBlank(t *testing.T) {
	var err error = Ent("")
	assert(t, "entity error", err.Error())

	entErr, ok := err.(EntityError)
	assert(t, true, ok)
	assertNotEqual(t, nil, entErr)
	assert(t, "", entErr.EntityIdentifier())
	assert(t, nil, Unwrap(err))
}

func TestEntityNoID(t *testing.T) {
	var err error = Ent("").Desc(ErrValueMalformed)
	assert(t, "entity malformed", err.Error())

	entErr, ok := err.(EntityError)
	assert(t, true, ok)
	assertNotEqual(t, nil, entErr)
	assert(t, "", entErr.EntityIdentifier())
	assert(t, nil, Unwrap(err))
	assert(t, ErrValueMalformed, UnwrapDescriptor(err))
}

func TestEntityWithFields(t *testing.T) {
	var err error = Ent("user").Fieldset(
		N("name").Desc(ErrValueEmpty),
		N("age").Desc(ErrValueUnspecified))
	assert(t, "user: name: empty, age: unspecified", err.Error())
}

func TestEntityWithFieldsNoName(t *testing.T) {
	var err error = Ent("").Fieldset(
		N("name").Desc(ErrValueEmpty),
		N("age").Desc(ErrValueUnspecified))
	assert(t, "entity: name: empty, age: unspecified", err.Error())
}

func TestEntNotFound(t *testing.T) {
	var fooNotFound error = Ent("foo").Wrap(ErrEntityNotFound)
	assert(t, "foo: not found", fooNotFound.Error())
	assert(t, true, IsEntityNotFoundError(fooNotFound))
}

func TestErrEntNotFound(t *testing.T) {
	var notFoundBare error = ErrEntityNotFound
	assert(t, false, IsEntityNotFoundError(notFoundBare))
	assert(t, false, IsArgumentError(notFoundBare))
	assert(t, false, IsCallError(notFoundBare))
}

func TestIsEntNotFoundErrorCustomNegative(t *testing.T) {
	var err error = &customEntError{entID: "foo"}
	if IsEntityNotFoundError(err) {
		t.Error(`IsEntNotFoundError(err)`)
	}
}

func TestEntDescMsg(t *testing.T) {
	var err error = Ent("foo").DescMsg("custom descriptor")
	assert(t, "foo: custom descriptor", err.Error())
	assert(t, "custom descriptor", UnwrapDescriptor(err).Error())
	assert(t, nil, Unwrap(err))
}

//----

type customEntError struct {
	entID string
}

var (
	_ EntityError = &customEntError{}
)

func (e *customEntError) EntityIdentifier() string { return e.entID }
func (e *customEntError) Error() string            { return "custom ent error" }
func (e *customEntError) CallError() CallError     { return e }
func (e *customEntError) Unwrap() error            { return nil }
func (e customEntError) FieldErrors() []NamedError { return nil }

func TestUnwrapEntityErrorSetNil(t *testing.T) {
	var err error
	if callErrors := UnwrapEntityErrorSet(err); len(callErrors) != 0 {
		t.Error("len(callErrors) != 0")
	}
}

func TestUnwrapEntityErrorSetWrongType(t *testing.T) {
	var err error = ErrValueMalformed
	if callErrors := UnwrapEntityErrorSet(err); len(callErrors) != 0 {
		t.Error("len(callErrors) != 0")
	}
}

func TestEntRewrapNil(t *testing.T) {
	var err error = Ent("foo").Rewrap(nil)
	assert(t, "foo", err.Error())
	assert(t, true, IsEntityError(err))
	assert(t, nil, UnwrapDescriptor(err))
	assert(t, nil, Unwrap(err))
}

func TestEntRewrapDesc(t *testing.T) {
	var err error = Ent("foo").Rewrap(ErrValueMalformed)
	assert(t, "foo: malformed", err.Error())
	assert(t, true, IsEntityError(err))
	assertNotEqual(t, nil, UnwrapDescriptor(err))
	assert(t, ErrValueMalformed, UnwrapDescriptor(err))
	assert(t, nil, Unwrap(err))
}

func TestEntRewrapDescWrapped(t *testing.T) {
	var err error = Ent("foo").Rewrap(Arg1().Desc(ErrValueMalformed).Wrap(Msg("bar")))
	assert(t, "foo: malformed: bar", err.Error())
	assert(t, true, IsEntityError(err))
	assertNotEqual(t, nil, UnwrapDescriptor(err))
	assert(t, ErrValueMalformed, UnwrapDescriptor(err))
	assertNotEqual(t, nil, Unwrap(err))
	assert(t, "bar", Unwrap(err).Error())
}

func TestEntRewrapRandom(t *testing.T) {
	var err error = Ent("foo").Rewrap(Msg("bar"))
	assert(t, "foo", err.Error())
	assert(t, true, IsEntityError(err))
	assert(t, nil, UnwrapDescriptor(err))
	assert(t, nil, Unwrap(err))
}

func TestEntRewrapWrappedNoDesc(t *testing.T) {
	var err error = Ent("foo").Rewrap(Arg1().Wrap(Msg("bar")))
	assert(t, "foo: bar", err.Error())
	assert(t, true, IsEntityError(err))
	assert(t, nil, UnwrapDescriptor(err))
	assertNotEqual(t, nil, Unwrap(err))
	assert(t, "bar", Unwrap(err).Error())
}

func TestEntRewrapFields(t *testing.T) {
	var err error = Ent("simple").Rewrap(Arg1().Fieldset(
		NamedValueUnsupported("foo"),
		N("bar").Desc(ErrValueMalformed),
	))
	assert(t, "simple: foo: unsupported, bar: malformed", err.Error())
	assert(t, true, IsEntityError(err))
	assert(t, nil, UnwrapDescriptor(err))
	assert(t, nil, Unwrap(err))
	assert(t, 2, len(UnwrapFieldErrors(err)))
}

//----

func TestEntSetEmpty(t *testing.T) {
	var err error = entErrorSet{}
	if err.Error() != "" {
		t.Errorf(`err.Error() != "" -- %q`, err.Error())
	}
	errSet := asErrorSet(err)
	if errSet == nil {
		t.Error("errSet == nil")
	}
	errs := errSet.Errors()
	if len(errs) != 0 {
		t.Error("len(errors) != 0")
	}
	errs = UnwrapErrorSet(err)
	if len(errs) != 0 {
		t.Error("len(errors) != 0")
	}
	entErrSet := asEntityErrorSet(err)
	if entErrSet == nil {
		t.Error("entErrSet == nil")
	}
	entErrors := entErrSet.EntityErrors()
	if len(entErrors) != 0 {
		t.Error("len(entErrors) != 0")
	}
	entErrors = UnwrapEntityErrorSet(err)
	if len(entErrors) != 0 {
		t.Error("len(entErrors) != 0")
	}
}

func TestEntSetSinge(t *testing.T) {
	var err error = EntSet(Ent("foo").Desc(ErrValueMalformed))
	if err.Error() != "foo: malformed" {
		t.Errorf(`err.Error() != "foo: malformed" -- %q`, err.Error())
	}
	errSet := asErrorSet(err)
	if errSet == nil {
		t.Error("errSet == nil")
	}
	errs := errSet.Errors()
	if len(errs) != 1 {
		t.Error("len(errors) != 1")
	}
	errs = UnwrapErrorSet(err)
	if len(errs) != 1 {
		t.Error("len(errors) != 1")
	}
	entErrSet := asEntityErrorSet(err)
	if entErrSet == nil {
		t.Error("entErrSet == nil")
	}
	entErrors := entErrSet.EntityErrors()
	if len(entErrors) != 1 {
		t.Error("len(entErrors) != 1")
	}
	entErrors = UnwrapEntityErrorSet(err)
	if len(entErrors) != 1 {
		t.Error("len(entErrors) != 1")
	}
}
