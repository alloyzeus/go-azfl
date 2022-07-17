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

	// Wrap returns a copy with wrapped error is set to detailingError.
	Wrap(detailingError error) OperationErrorBuilder
}

func Op(operationName string) OperationErrorBuilder {
	return &opError{operationName: operationName}
}

type opError struct {
	operationName string
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
	causeStr := errorString(e.wrapped)
	if e.operationName != "" {
		if causeStr != "" {
			return e.operationName + ": " + causeStr
		}
		return e.operationName + " error"
	}
	if causeStr != "" {
		return "operation error: " + causeStr
	}
	return "operation error"
}

func (e opError) Wrap(detailingError error) OperationErrorBuilder {
	e.wrapped = detailingError
	return &e
}