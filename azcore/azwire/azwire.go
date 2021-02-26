package azwire

// DataType represents a type supported by azwire.
type DataType uint8

// Supported data types.
//
// MSB is reserved
const (
	// DataTypeUnspecified is used for the zero value of DataType.
	DataTypeUnspecified DataType = 0b0
	// DataTypeArray is used to indicate an array data type.
	DataTypeArray DataType = 0b01000000

	DataTypeInt16 DataType = 0b00010010
	DataTypeInt32 DataType = 0b00010011
	DataTypeInt64 DataType = 0b00010100
)

// DataTypeFromByte parses a byte to its respective DataType.
func DataTypeFromByte(b byte) (DataType, error) {
	//TODO: check for invalid value
	return DataType(b), nil
}

// Byte returns a representation of DataType as a byte. The byte then can
// be used as the byte mark in the resulting encoded data.
func (typ DataType) Byte() byte { return byte(typ) }

// Marshalable provides a contract for objects with can be marshaled as
// azwire. If the object is part of other object, i.e., a
// field, see FieldMarshalable.
type Marshalable interface {
	// AZWire returns an azwire-encoded representation of the instance as
	// a top-level object.
	//
	// The resulting azwire-encoded data is prepended with the data type
	// marker byte.
	AZWire() []byte
}

// FieldMarshalable provides a contract for objects which can be marshaled
// as a structure field into azwire.
type FieldMarshalable interface {
	// AZWireField returns an azwire-encoded representation of the instance.
	// The data type is returned along with the data bytes instead of included
	// in them as this method was designed for constructing an azwire-encoded
	// structure where this object is part of.
	AZWireField() ([]byte, DataType)
}
