package errors

// Something that has more semantic than "Wrap"
// See https://github.com/alloyzeus/go-azfl/issues/8 for the discussions on
// what we want to implement.

type OperationError interface {
	Unwrappable

	// Returns the name of the operation.
	OperationName() string
}

type OperationErrorBuilder interface {
	OperationError

	// Hint provides a clue for the developers on how to fix the error.
	Hint(hintText string) OperationErrorBuilder

	Params(params ...NamedError) OperationErrorBuilder

	// Wrap returns a copy with wrapped error is set to detailingError.
	Wrap(detailingError error) OperationErrorBuilder
}

func Op(operationName string, detailingError error) OperationErrorBuilder {
	return &opError{operationName: operationName, wrapped: detailingError}
}

type opError struct {
	operationName string
	hintText      string
	params        []NamedError
	wrapped       error
}

var (
	_ error                 = &opError{}
	_ Unwrappable           = &opError{}
	_ OperationError        = &opError{}
	_ OperationErrorBuilder = &opError{}
)

func (e opError) OperationName() string { return e.operationName }

func (e opError) Unwrap() error { return e.wrapped }

func (e *opError) Error() string {
	suffix := namedSetToString(e.params)
	if suffix != "" {
		suffix = ". " + suffix
	}
	if e.hintText != "" {
		suffix = suffix + ". " + e.hintText
	}
	causeStr := errorString(e.wrapped)

	if e.operationName != "" {
		if causeStr != "" {
			return e.operationName + ": " + causeStr + suffix
		}
		if suffix != "" {
			return e.operationName + " error" + suffix
		}
		return e.operationName + " error"
	}
	if causeStr != "" {
		return "operation " + causeStr + suffix
	}
	if suffix != "" {
		return "operation error" + suffix
	}
	return "operation error"
}

func (e opError) Hint(hintText string) OperationErrorBuilder {
	e.hintText = hintText
	return &e
}

func (e opError) Params(params ...NamedError) OperationErrorBuilder {
	e.params = copyNamedSet(params)
	return &e
}

func (e opError) Wrap(detailingError error) OperationErrorBuilder {
	e.wrapped = detailingError
	return &e
}
