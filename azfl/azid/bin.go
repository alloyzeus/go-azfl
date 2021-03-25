package azid

// BinMarshalable is an interface definition for objects which able to
// provide an azid-bin representation of itself.
//
// This interface is for top-level object. If the object is part of other
// object, i.e., a field, see BinFieldMarshalable.
type BinMarshalable interface {
	// AZIDBin returns an azid-bin representation of the instance.
	AZIDBin() []byte
}

// BinUnmarshalable is and interface definition for objects which able
// to load an azid-bin representation into itself.
type BinUnmarshalable interface {
	UnmarshalAZIDBin(b []byte) (readLen int, err error)
}

// BinFieldMarshalable is an interface definition for objects which
// able to provide an azid-bin representation of itself as a field
// of other azid object.
type BinFieldMarshalable interface {
	// AZIDBinField returns an azid-bin representation of the instance
	// which was designed to be part of the larger object.
	// The data type is returned along with the data bytes instead of included
	// in them as this method was designed for constructing an azid-bin
	// structure where this object is part of.
	AZIDBinField() ([]byte, BinDataType)
}

// BinFieldUnmarshalable is and interface definition for objects which able
// to load an azid-bin representation into itself.
type BinFieldUnmarshalable interface {
	UnmarshalAZIDBinField(b []byte, typeHint BinDataType) (readLen int, err error)
}

//region BinDataType

// BinDataType represents a type supported by azid-bin.
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
