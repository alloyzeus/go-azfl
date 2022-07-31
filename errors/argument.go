package errors

//TODO: ArgSet

// ArgumentError abstracts all errors which was caused by error in
// one of the arguments in a function call. This class of error
// has the similar concept as 4xx status codes in HTTP, and thus can
// be mapped into one of these code when used in HTTP request handlers.
type ArgumentError interface {
	CallError
	Unwrappable

	// ArgumentName returns the name of the offending argument. It might
	// be empty if there's only one argument.
	ArgumentName() string

	Descriptor() ErrorDescriptor
	FieldErrors() []NamedError
}

type ArgumentErrorBuilder interface {
	ArgumentError

	// Desc returns a copy with descriptor is set to desc.
	Desc(desc ErrorDescriptor) ArgumentErrorBuilder

	// DescMsg sets the descriptor with the provided string. For the best
	// experience, descMsg should be defined as a constant so that the error
	// users could use it to identify an error. For non-constant descriptor
	// use the Wrap method.
	DescMsg(descMsg string) ArgumentErrorBuilder

	// Hint provides a clue for the developers on how to fix the error.
	Hint(hintText string) ArgumentErrorBuilder

	Fieldset(fields ...NamedError) ArgumentErrorBuilder

	// Rewrap collects descriptor, wrapped, and fields from err and include
	// them into the new error.
	Rewrap(err error) ArgumentErrorBuilder

	// Wrap returns a copy with wrapped error is set to detailingError.
	Wrap(detailingError error) ArgumentErrorBuilder
}

func IsArgumentError(err error) bool {
	_, ok := err.(ArgumentError)
	return ok
}

// AsArgumentError returns non-nil if err is indeed an ArgumentError.
func AsArgumentError(err error) ArgumentError {
	if e, ok := err.(ArgumentError); ok {
		return e
	}
	return nil
}

func Arg(argName string) ArgumentErrorBuilder {
	return &argumentError{
		argName: argName,
	}
}

// Arg1 is used when there's only one argument for a function.
func Arg1() ArgumentErrorBuilder {
	return &argumentError{
		argName: "",
	}
}

func ArgFields(argName string, fields ...NamedError) ArgumentErrorBuilder {
	return &argumentError{
		argName: argName,
		fields:  fields,
	}
}

func ArgMsg(argName, errMsg string, fields ...NamedError) ArgumentErrorBuilder {
	return &argumentError{
		argName: argName,
		wrapped: Msg(errMsg),
		fields:  fields,
	}
}

func ArgWrap(argName, contextMessage string, err error, fields ...NamedError) ArgumentErrorBuilder {
	return &argumentError{
		argName: argName,
		wrapped: Wrap(contextMessage, err),
		fields:  fields,
	}
}

// ArgUnspecified describes that argument with name argName is unspecified.
func ArgUnspecified(argName string) ArgumentErrorBuilder {
	return &argumentError{
		argName:    argName,
		descriptor: ErrValueUnspecified,
	}
}

func IsArgumentUnspecifiedError(err error) bool {
	if !IsArgumentError(err) {
		return false
	}
	if desc := UnwrapDescriptor(err); desc != nil {
		return desc == ErrValueUnspecified
	}
	return false
}

// ArgValueUnsupported creates an ArgumentError with name is set to the value
// of argName and descriptor is set to ErrValueUnsupported.
func ArgValueUnsupported(argName string) ArgumentErrorBuilder {
	return &argumentError{
		argName:    argName,
		descriptor: ErrValueUnsupported,
	}
}

// IsArgumentUnspecified checks if an error describes about unspecifity of an argument.
//
//TODO: ArgSet
func IsArgumentUnspecified(err error, argName string) bool {
	if err == nil {
		return false
	}
	argErr, ok := err.(ArgumentError)
	if !ok {
		return false
	}
	if argErr.ArgumentName() != argName {
		return false
	}
	if desc := UnwrapDescriptor(err); desc != nil {
		return desc == ErrValueUnspecified
	}
	return false
}

func IsArgUnspecified(err error, argName string) bool {
	if err == nil {
		return false
	}
	argErr, ok := err.(ArgumentError)
	if !ok {
		return false
	}
	if argErr.ArgumentName() != argName {
		return false
	}
	if desc := UnwrapDescriptor(err); desc != nil {
		return desc == ErrValueUnspecified
	}
	return false
}

type argumentError struct {
	argName    string
	descriptor ErrorDescriptor
	hintText   string
	fields     []NamedError
	wrapped    error
}

var (
	_ error                = &argumentError{}
	_ Unwrappable          = &argumentError{}
	_ CallError            = &argumentError{}
	_ ArgumentError        = &argumentError{}
	_ ArgumentErrorBuilder = &argumentError{}
	_ hasDescriptor        = &argumentError{}
	_ hasFieldErrors       = &argumentError{}
)

func (e *argumentError) ArgumentName() string {
	return e.argName
}

func (e *argumentError) CallError() CallError { return e }

func (e *argumentError) Error() string {
	suffix := namedSetToString(e.fields)
	if suffix != "" {
		suffix = ": " + suffix
	}
	if e.hintText != "" {
		suffix = suffix + ". " + e.hintText
	}
	var descStr string
	if e.descriptor != nil {
		descStr = e.descriptor.Error()
	}
	causeStr := errorString(e.wrapped)
	if causeStr == "" {
		causeStr = descStr
	} else if descStr != "" {
		causeStr = descStr + ": " + causeStr
	}

	if e.argName != "" {
		if causeStr != "" {
			return "arg " + e.argName + ": " + causeStr + suffix
		}
		return "arg " + e.argName + suffix
	}
	if causeStr != "" {
		return "arg " + causeStr + suffix
	}
	if suffix != "" {
		return "arg" + suffix
	}
	return "arg error"
}

func (e argumentError) Descriptor() ErrorDescriptor { return e.descriptor }
func (e argumentError) Unwrap() error               { return e.wrapped }
func (e argumentError) FieldErrors() []NamedError   { return copyNamedSet(e.fields) }

func (e argumentError) Desc(desc ErrorDescriptor) ArgumentErrorBuilder {
	e.descriptor = desc
	return &e
}

func (e argumentError) DescMsg(descMsg string) ArgumentErrorBuilder {
	e.descriptor = constantErrorDescriptor(descMsg)
	return &e
}

func (e argumentError) Hint(hintText string) ArgumentErrorBuilder {
	e.hintText = hintText
	return &e
}

func (e argumentError) Fieldset(fields ...NamedError) ArgumentErrorBuilder {
	e.fields = copyNamedSet(fields)
	return &e
}

func (e argumentError) Rewrap(err error) ArgumentErrorBuilder {
	if err != nil {
		if descErr, _ := err.(ErrorDescriptor); descErr != nil {
			e.descriptor = descErr
			e.wrapped = nil
			e.fields = nil
		} else {
			e.descriptor = UnwrapDescriptor(err)
			e.wrapped = Unwrap(err)
			e.fields = UnwrapFieldErrors(err)
		}
	}
	return &e
}

func (e argumentError) Wrap(detailingError error) ArgumentErrorBuilder {
	e.wrapped = detailingError
	return &e
}

type ArgumentErrorChecker struct {
	errToCheck ArgumentError
	truthState bool
}

func ArgumentErrorCheck(err error) ArgumentErrorChecker {
	argErr, ok := err.(ArgumentError)
	return ArgumentErrorChecker{
		errToCheck: argErr,
		truthState: ok,
	}
}

func (checker ArgumentErrorChecker) IsTrue() bool { return checker.truthState }

func (checker ArgumentErrorChecker) HasName(argName string) ArgumentErrorChecker {
	if err := checker.errToCheck; err == nil || err.ArgumentName() != argName {
		checker.truthState = false
	}
	return checker
}

func (checker ArgumentErrorChecker) HasDesc(desc ErrorDescriptor) ArgumentErrorChecker {
	if !HasDescriptor(checker.errToCheck, desc) {
		checker.truthState = false
	}
	return checker
}

func (checker ArgumentErrorChecker) HasWrapped(err error) ArgumentErrorChecker {
	if !Is(err, Unwrap(checker.errToCheck)) {
		checker.truthState = false
	}
	return checker
}
