package azer

import "github.com/alloyzeus/go-azcore/azcore/errors"

// BinMarshalable is an interface definition for objects which able to
// provide an azer-bin-encoded representation of itself.
//
// This interface is for top-level object. If the object is part of other
// object, i.e., a field, see BinFieldMarshalable.
type BinMarshalable interface {
	// AZERBin returns an azer-bin-encoded representation of the instance.
	AZERBin() []byte
}

// BinUnmarshalable is and interface definition for objects which able
// to load an azer-bin-encoded representation into itself.
type BinUnmarshalable interface {
	UnmarshalAZERBin(b []byte) (readLen int, err error)
}

// BinFieldMarshalable is an interface definition for objects which
// able to provide an azer-bin-encoded representation of itself as a field
// of other azer object.
type BinFieldMarshalable interface {
	// AZERBinField returns an azer-bin-encoded representation of the instance
	// which was designed to be part of the larger object.
	// The data type is returned along with the data bytes instead of included
	// in them as this method was designed for constructing an azer-bin-encoded
	// structure where this object is part of.
	AZERBinField() ([]byte, BinDataType)
}

// BinFieldUnmarshalable is and interface definition for objects which able
// to load an azer-bin-encoded representation into itself.
type BinFieldUnmarshalable interface {
	UnmarshalAZERBinField(b []byte, typeHint BinDataType) (readLen int, err error)
}

//region BinDataType

// BinDataType represents a type supported by azer-bin.
type BinDataType uint8

// Supported data types.
//
// MSB is reserved
const (
	// BinDataTypeUnspecified is used for the zero value of BinDataType.
	BinDataTypeUnspecified BinDataType = 0b0
	// BinDataTypeArray is used to indicate an array data type.
	BinDataTypeArray BinDataType = 0b01000000

	BinDataTypeInt16 BinDataType = 0b00010010
	BinDataTypeInt32 BinDataType = 0b00010011
	BinDataTypeInt64 BinDataType = 0b00010100
)

// BinDataTypeFromByte parses a byte to its respective BinDataType.
func BinDataTypeFromByte(b byte) (BinDataType, error) {
	//TODO: accept only valid value
	return BinDataType(b), nil
}

// Byte returns a representation of BinDataType as a byte. The byte then can
// be used as the byte mark in the resulting encoded data.
func (typ BinDataType) Byte() byte { return byte(typ) }

//endregion

// BinError abstracts all errors emitted by azer-bin operations.
type BinError interface {
	Error
}

// BinUnmarshalError abstracts all errors emitted by azer-bin unmarshalling
// operations.
type BinUnmarshalError interface {
	BinError
}

// BinUnmarshalArgumentError is a specialization of BinUnmarshalError which
// indicates that the error is in the one of the arguments.
type BinUnmarshalArgumentError interface {
	BinUnmarshalError
	errors.ArgumentError
}

// BinUnmarshalArgumentErrorMsg returns an instance of
// BinUnmarshalArgumentError with simple error message.
func BinUnmarshalArgumentErrorMsg(
	argName string,
	errMsg string,
	fields ...errors.EntityError,
) BinUnmarshalArgumentError {
	return binUnmarshalArgumentError{
		argName: argName,
		err:     errors.Msg(errMsg),
		fields:  fields,
	}
}

// BinUnmarshalArgumentErrorWrap wraps an error.
func BinUnmarshalArgumentErrorWrap(
	argName string,
	contextMessage string,
	err error,
	fields ...errors.EntityError,
) BinUnmarshalArgumentError {
	return binUnmarshalArgumentError{
		argName: argName,
		err:     errors.Wrap(contextMessage, err),
		fields:  fields,
	}
}

type binUnmarshalArgumentError struct {
	argName string
	err     error
	fields  []errors.EntityError
}

var (
	_ error                = binUnmarshalArgumentError{}
	_ errors.Unwrappable   = binUnmarshalArgumentError{}
	_ errors.CallError     = binUnmarshalArgumentError{}
	_ errors.EntityError   = binUnmarshalArgumentError{}
	_ errors.ArgumentError = binUnmarshalArgumentError{}
)

func (e binUnmarshalArgumentError) EntityIdentifier() string { return e.argName }

func (e binUnmarshalArgumentError) ArgumentName() string {
	return e.argName
}

func (binUnmarshalArgumentError) CallError() {}

func (e binUnmarshalArgumentError) Error() string {
	if e.argName != "" {
		if errMsg := e.err.Error(); errMsg != "" {
			return "arg " + e.argName + ": " + errMsg
		}
		return "arg " + e.argName + " invalid"
	}
	if errMsg := e.err.Error(); errMsg != "" {
		return "arg " + errMsg
	}
	return "arg invalid"
}

func (e binUnmarshalArgumentError) Unwrap() error {
	return e.err
}
