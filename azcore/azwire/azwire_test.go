package azwire_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/alloyzeus/go-azcore/azcore/azwire"
	"github.com/alloyzeus/go-azcore/azcore/errors"
)

type int32ID int32

var _ azwire.FieldMarshalable = int32ID(0)

func (id int32ID) AZWireField() ([]byte, azwire.DataType) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(id))
	return b, azwire.DataTypeInt32
}

func int32IDFromAZWireField(
	b []byte, typeHint azwire.DataType,
) (id int32ID, readLen int, err error) {
	if typeHint != azwire.DataTypeUnspecified && typeHint != azwire.DataTypeInt32 {
		return int32ID(0), 0,
			errors.Msg("unsupported parsing from the buffer with the specified type")
	}
	i := binary.BigEndian.Uint32(b)
	return int32ID(i), 4, nil
}

type int32RefKey int32ID

var _ azwire.Marshalable = int32RefKey(0)
var _ azwire.FieldMarshalable = int32RefKey(0)

func int32RefKeyFromAZWireField(
	b []byte, typeHint azwire.DataType,
) (refKey int32RefKey, readLen int, err error) {
	id, n, err := int32IDFromAZWireField(b, typeHint)
	if err != nil {
		return int32RefKey(0), n, err
	}
	return int32RefKey(id), n, nil
}

func (refKey int32RefKey) AZWire() []byte {
	b := make([]byte, 5)
	b[0] = azwire.DataTypeInt32.Byte()
	binary.BigEndian.PutUint32(b[1:], uint32(refKey))
	return b
}

func (refKey int32RefKey) AZWireField() ([]byte, azwire.DataType) {
	return int32ID(refKey).AZWireField()
}

type adjunctID int16

func adjunctIDFromAZWireField(
	b []byte, typeHint azwire.DataType,
) (id adjunctID, readLen int, err error) {
	if typeHint != azwire.DataTypeUnspecified && typeHint != azwire.DataTypeInt16 {
		return adjunctID(0), 0,
			errors.Msg("unsupported parsing from the buffer with the specified type")
	}
	i := binary.BigEndian.Uint16(b)
	return adjunctID(i), 2, nil
}

func (id adjunctID) AZWireField() ([]byte, azwire.DataType) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(id))
	return b, azwire.DataTypeInt16
}

type adjunctRefKey struct {
	parent int32RefKey
	id     adjunctID
}

func adjunctRefKeyFromAZWire(
	b []byte,
) (refKey adjunctRefKey, readLen int, err error) {
	typ, err := azwire.DataTypeFromByte(b[0])
	if err != nil {
		return adjunctRefKey{}, 0, err
	}
	if typ != azwire.DataTypeArray {
		return adjunctRefKey{}, 0,
			errors.Msg("unsupported parsing from the buffer with the specified type")
	}

	arrayLen := int(b[1])
	if arrayLen != 2 {
		return adjunctRefKey{}, 0,
			errors.Msg("unexpected number of array length")
	}

	parentType, err := azwire.DataTypeFromByte(b[2])
	if err != nil {
		return adjunctRefKey{}, 0, err
	}
	parentRefKey, _, err := int32RefKeyFromAZWireField(b[4:], parentType)
	if err != nil {
		return adjunctRefKey{}, 0,
			errors.Msg("unable to parse")
	}

	idType, err := azwire.DataTypeFromByte(b[3])
	if err != nil {
		return adjunctRefKey{}, 0, err
	}
	id, _, err := adjunctIDFromAZWireField(b[8:], idType)
	if err != nil {
		return adjunctRefKey{}, 0,
			errors.Msg("unable to parse")
	}

	return adjunctRefKey{parentRefKey, id}, 10, nil
}

func (refKey adjunctRefKey) AZWire() []byte {
	var fieldTypes []byte
	var fieldData []byte

	b, t := refKey.parent.AZWireField()
	fieldTypes = append(fieldTypes, t.Byte())
	fieldData = append(fieldData, b...)

	b, t = refKey.id.AZWireField()
	fieldTypes = append(fieldTypes, t.Byte())
	fieldData = append(fieldData, b...)

	var out = []byte{azwire.DataTypeArray.Byte(), byte(len(fieldTypes))}
	out = append(out, fieldTypes...)
	out = append(out, fieldData...)
	return out
}

func TestEncodeField(t *testing.T) {
	testCases := []struct {
		in      int32ID
		outData []byte
		outType azwire.DataType
	}{
		{int32ID(0), []byte{0, 0, 0, 0}, azwire.DataTypeInt32},
		{int32ID(1), []byte{0, 0, 0, 1}, azwire.DataTypeInt32},
		{int32ID(1 << 24), []byte{1, 0, 0, 0}, azwire.DataTypeInt32},
	}

	for _, testCase := range testCases {
		outData, outType := testCase.in.AZWireField()
		if !bytes.Equal(outData, testCase.outData) {
			t.Errorf("Expected: %#v, actual: %#v", testCase.outData, outData)
		}
		if outType != testCase.outType {
			t.Errorf("Expected: %#v, actual: %#v", testCase.outType, outType)
		}
	}
}

func TestEncodeRefKey(t *testing.T) {
	testCases := []struct {
		in      int32RefKey
		outData []byte
	}{
		{int32RefKey(0), []byte{0x13, 0, 0, 0, 0}},
		{int32RefKey(1), []byte{0x13, 0, 0, 0, 1}},
		{int32RefKey(1 << 24), []byte{0x13, 1, 0, 0, 0}},
	}

	for _, testCase := range testCases {
		outData := testCase.in.AZWire()
		if !bytes.Equal(outData, testCase.outData) {
			t.Errorf("Expected: %#v, actual: %#v", testCase.outData, outData)
		}
	}
}

func TestEncodeAdjunct(t *testing.T) {
	testCases := []struct {
		in      adjunctRefKey
		outData []byte
	}{
		{adjunctRefKey{int32RefKey(0), adjunctID(0)},
			[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 0, 0, 0}},
		{adjunctRefKey{int32RefKey(0), adjunctID(1)},
			[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 0, 0, 1}},
		{adjunctRefKey{int32RefKey(1), adjunctID(0)},
			[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 1, 0, 0}},
	}

	for _, testCase := range testCases {
		outData := testCase.in.AZWire()
		if !bytes.Equal(outData, testCase.outData) {
			t.Errorf("Expected: %#v, actual: %#v", testCase.outData, outData)
		}
	}
}

func TestDecodeToAdjunct(t *testing.T) {
	testCases := []struct {
		in      []byte
		out     adjunctRefKey
		readLen int
		err     error
	}{
		{[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 0, 0, 0},
			adjunctRefKey{int32RefKey(0), adjunctID(0)}, 10, nil},
		{[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 0, 0, 1},
			adjunctRefKey{int32RefKey(0), adjunctID(1)}, 10, nil},
		{[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 1, 0, 0},
			adjunctRefKey{int32RefKey(1), adjunctID(0)}, 10, nil},
	}

	for _, testCase := range testCases {
		refKey, readLen, err := adjunctRefKeyFromAZWire(testCase.in)
		if err != testCase.err {
			t.Errorf("Expected: %#v, actual: %#v", testCase.err, err)
		}
		if readLen != testCase.readLen || readLen != len(testCase.in) {
			t.Errorf("Expected: %#v, actual: %#v", testCase.readLen, readLen)
		}
		if refKey.parent != testCase.out.parent || refKey.id != testCase.out.id {
			t.Errorf("Expected: %#v, actual: %#v", testCase.out, refKey)
		}
	}
}
