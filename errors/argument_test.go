package errors

import "testing"

func TestArgEmpty(t *testing.T) {
	var err error = Arg("")
	assert(t, "arg error", err.Error())

	inner := Unwrap(err)
	assert(t, nil, inner)
	assert(t, false, IsArgumentUnspecifiedError(err))
}

func TestArgFields(t *testing.T) {
	var err error = Arg("foo").Fieldset(
		N("name").Desc(ErrValueEmpty), N("bar").Desc(ErrValueUnspecified))
	assert(t, "arg foo: name: empty, bar: unspecified", err.Error())
}

func TestArgFieldsNoName(t *testing.T) {
	var err error = Arg("").Fieldset(
		N("name").Desc(ErrValueEmpty), N("bar").Desc(ErrValueUnspecified))
	assert(t, "arg: name: empty, bar: unspecified", err.Error())
}

func TestArg1(t *testing.T) {
	var err error = Arg1().Desc(ErrValueUnspecified)
	assert(t, "arg unspecified", err.Error())
}

func TestArgDescWrap(t *testing.T) {
	inner := Msg("inner")
	var err error = Arg("foo").Desc(ErrValueMalformed).Wrap(inner)
	assert(t, "arg foo: malformed: inner", err.Error())
	assert(t, true, ArgumentErrorCheck(err).IsTrue())
	assert(t, true, ArgumentErrorCheck(err).HasName("foo").IsTrue())
	assert(t, false, ArgumentErrorCheck(err).HasName("bar").IsTrue())
	assert(t, true, ArgumentErrorCheck(err).HasName("foo").HasDesc(ErrValueMalformed).IsTrue())
	assert(t, false, ArgumentErrorCheck(err).HasName("foo").HasDesc(ErrValueEmpty).IsTrue())
	assert(t, false, ArgumentErrorCheck(err).HasName("bar").HasDesc(ErrValueMalformed).IsTrue())
	assert(t, false, ArgumentErrorCheck(err).HasName("bar").HasDesc(ErrValueEmpty).IsTrue())
	assert(t, true, ArgumentErrorCheck(err).HasName("foo").HasDesc(ErrValueMalformed).HasWrapped(inner).IsTrue())
	assert(t, false, ArgumentErrorCheck(err).HasWrapped(nil).IsTrue())
	assert(t, false, ArgumentErrorCheck(inner).IsTrue())

	wrapped := Unwrap(err)
	assertNotEqual(t, nil, wrapped)
	assert(t, inner, wrapped)

	assert(t, true, HasDescriptor(err, ErrValueMalformed))

	desc := UnwrapDescriptor(err)
	assertNotEqual(t, nil, desc)
	assert(t, ErrValueMalformed, desc)
}

func TestArgDescMsg(t *testing.T) {
	const customDesc = "custom descriptor"
	var err error = Arg("saturn").DescMsg(customDesc)
	assert(t, "arg saturn: custom descriptor", err.Error())
	assert(t, true, HasDescriptorText(err, customDesc))

	desc := UnwrapDescriptor(err)
	assertNotEqual(t, nil, desc)
	assert(t, customDesc, desc.Error())
}

func TestAsArgumentError(t *testing.T) {
	var err error = Arg1()
	assert(t, err, AsArgumentError(err))
}

func TestAsArgumentErrorNil(t *testing.T) {
	var err error
	assert(t, nil, AsArgumentError(err))
}

func TestAsArgumentErrorNegative(t *testing.T) {
	var err error = N("")
	assert(t, nil, AsArgumentError(err))
}

func TestArgUnspecifiedEmpty(t *testing.T) {
	var err error = ArgUnspecified("")

	assert(t, "arg unspecified", err.Error())
	assert(t, true, IsArgumentUnspecifiedError(err))
	assert(t, ErrValueUnspecified, UnwrapDescriptor(err))

	argErr, ok := err.(ArgumentError)
	assert(t, true, ok)
	assertNotEqual(t, nil, argErr)
	assert(t, argErr, argErr.CallError())
	assert(t, "", argErr.ArgumentName())

	argErrB := AsArgumentError(err)
	assertNotEqual(t, nil, argErrB)
	assert(t, true, argErr == argErrB)
	assert(t, argErr, argErrB)

	wrapped := Unwrap(err)
	assert(t, nil, wrapped)

	desc := UnwrapDescriptor(err)
	assertNotEqual(t, nil, desc)
	assert(t, ErrValueUnspecified, desc)
}

func TestArgUnspecifiedFoo(t *testing.T) {
	var err error = ArgUnspecified("foo")
	assert(t, "arg foo: unspecified", err.Error())
	assert(t, true, IsArgumentUnspecifiedError(err))
	assert(t, true, IsArgumentUnspecified(err, "foo"))
	assert(t, false, IsArgumentUnspecified(err, "bar"))

	argErr, ok := err.(ArgumentError)
	assert(t, true, ok)
	assertNotEqual(t, nil, argErr)
	assert(t, "foo", argErr.ArgumentName())
	assert(t, true, IsArgumentUnspecifiedError(err))
	assert(t, ErrValueUnspecified, UnwrapDescriptor(err))

	wrapped := Unwrap(err)
	assert(t, nil, wrapped)

	desc := UnwrapDescriptor(err)
	assertNotEqual(t, nil, desc)
	assert(t, ErrValueUnspecified, desc)
}

func TestIsArgUnspecifiedErrorNil(t *testing.T) {
	var err error
	assert(t, false, IsArgumentUnspecifiedError(err))
	assert(t, false, IsArgumentUnspecified(err, "foo"))
}

func TestIsArgUnspecifiedNegative(t *testing.T) {
	var err error = ErrValueInvalid
	assert(t, false, IsArgumentUnspecified(err, "foo"))
}

func TestIsArgUnspecifiedWrongArgName(t *testing.T) {
	var err error = ArgUnspecified("foo")
	assert(t, false, IsArgumentUnspecified(err, "bar"))
}

func TestIsArgUnspecifiedCustomStruct(t *testing.T) {
	var err error = &customArgError{argName: "foo"}
	assert(t, false, IsArgumentUnspecified(err, "foo"))
}

func TestArgValueUnsupportedNoName(t *testing.T) {
	var err error = ArgValueUnsupported("")

	//TODO: should be "arg value unsupported"
	assert(t, "arg unsupported", err.Error())
	assert(t, nil, Unwrap(err))
	assert(t, ErrValueUnsupported, UnwrapDescriptor(err))

	argErr, ok := err.(ArgumentError)
	assert(t, true, ok)
	assertNotEqual(t, nil, argErr)
	assert(t, "", argErr.ArgumentName())
}

func TestArgValueUnsupportedFoo(t *testing.T) {
	var err error = ArgValueUnsupported("foo")
	assert(t, "arg foo: unsupported", err.Error())
	// assert(t, true, IsArgumentUnspecifiedError(err))
	// assert(t, true, IsArgumentUnspecified(err, "foo"))
	// assert(t, false, IsArgumentUnspecified(err, "bar"))

	argErr, ok := err.(ArgumentError)
	assert(t, true, ok)
	assertNotEqual(t, nil, argErr)
	assert(t, "foo", argErr.ArgumentName())
	// assert(t, true, IsArgumentUnspecifiedError(err))
	assert(t, ErrValueUnsupported, UnwrapDescriptor(err))

	wrapped := Unwrap(err)
	assert(t, nil, wrapped)

	desc := UnwrapDescriptor(err)
	assertNotEqual(t, nil, desc)
	assert(t, ErrValueUnsupported, desc)
}

func TestArgRewrapNil(t *testing.T) {
	var err error = Arg("foo").Rewrap(nil)
	assert(t, "arg foo", err.Error())
	assert(t, true, IsArgumentError(err))
	assert(t, false, IsEntityError(err))
	assert(t, nil, UnwrapDescriptor(err))
	assert(t, nil, Unwrap(err))
}

func TestArgRewrapDesc(t *testing.T) {
	var err error = Arg("foo").Rewrap(ErrValueMalformed)
	assert(t, "arg foo: malformed", err.Error())
	assert(t, true, IsArgumentError(err))
	assert(t, false, IsEntityError(err))
	assertNotEqual(t, nil, UnwrapDescriptor(err))
	assert(t, ErrValueMalformed, UnwrapDescriptor(err))
	assert(t, nil, Unwrap(err))
}

func TestArgRewrapDescWrapped(t *testing.T) {
	var err error = Arg("foo").Rewrap(Arg1().Desc(ErrValueMalformed).Wrap(Msg("bar")))
	assert(t, "arg foo: malformed: bar", err.Error())
	assert(t, true, IsArgumentError(err))
	assert(t, false, IsEntityError(err))
	assertNotEqual(t, nil, UnwrapDescriptor(err))
	assert(t, ErrValueMalformed, UnwrapDescriptor(err))
	assertNotEqual(t, nil, Unwrap(err))
	assert(t, "bar", Unwrap(err).Error())
}

func TestArgRewrapRandom(t *testing.T) {
	var err error = Arg("foo").Rewrap(Msg("bar"))
	assert(t, "arg foo", err.Error())
	assert(t, true, IsArgumentError(err))
	assert(t, false, IsEntityError(err))
	assert(t, nil, UnwrapDescriptor(err))
	assert(t, nil, Unwrap(err))
}

func TestArgRewrapWrappedNoDesc(t *testing.T) {
	var err error = Arg("foo").Rewrap(Arg1().Wrap(Msg("bar")))
	assert(t, "arg foo: bar", err.Error())
	assert(t, true, IsArgumentError(err))
	assert(t, false, IsEntityError(err))
	assert(t, nil, UnwrapDescriptor(err))
	assertNotEqual(t, nil, Unwrap(err))
	assert(t, "bar", Unwrap(err).Error())
}

func TestArgRewrapFields(t *testing.T) {
	var err error = Arg("simple").Rewrap(Arg1().Fieldset(
		NamedValueUnsupported("foo"),
		N("bar").Desc(ErrValueMalformed),
	))
	assert(t, "arg simple: foo: unsupported, bar: malformed", err.Error())
	assert(t, true, IsArgumentError(err))
	assert(t, false, IsEntityError(err))
	assert(t, nil, UnwrapDescriptor(err))
	assert(t, nil, Unwrap(err))
	assert(t, 2, len(UnwrapFieldErrors(err)))
	assert(t, true, ArgumentErrorCheck(err).IsTrue())
	assert(t, true, ArgumentErrorCheck(err).HasName("simple").IsTrue())
	assert(t, true, ArgumentErrorCheck(err).HasName("simple").HasDesc(nil).IsTrue())
	assert(t, true, ArgumentErrorCheck(err).HasName("simple").HasDesc(nil).HasWrapped(nil).IsTrue())
}

func TestArgHint(t *testing.T) {
	var err error = Arg("color").Hint("Check the color of the sky right now")

	assert(t, "arg color. Check the color of the sky right now", err.Error())
}

func TestArgHintDesc(t *testing.T) {
	var err error = ArgUnspecified("color").Hint("Check the color of the ocean right now")

	assert(t, "arg color: unspecified. Check the color of the ocean right now", err.Error())
}

func TestArgHintWrap(t *testing.T) {
	var err error = Arg("color").Wrap(Msg("hex invalid")).Hint("Check the color of your right iris")

	assert(t, "arg color: hex invalid. Check the color of your right iris", err.Error())
}

// ----

type customArgError struct {
	argName string
}

var (
	_ ArgumentError = &customArgError{}
)

func (e *customArgError) ArgumentName() string     { return e.argName }
func (e *customArgError) Error() string            { return "custom arg error" }
func (e *customArgError) CallError() CallError     { return e }
func (e *customArgError) Unwrap() error            { return nil }
func (customArgError) Descriptor() ErrorDescriptor { return nil }
func (e customArgError) FieldErrors() []NamedError { return nil }
