package azid_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/alloyzeus/go-azfl/azfl/azid"
	"github.com/alloyzeus/go-azfl/azfl/errors"
)

type int32IDNum int32

var _ azid.BinFieldMarshalable = int32IDNum(0)

func (id int32IDNum) AZIDBinField() ([]byte, azid.BinDataType) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(id))
	return b, azid.BinDataTypeInt32
}

func (id *int32IDNum) UnmarshalAZIDBinField(b []byte, typeHint azid.BinDataType) (readLen int, err error) {
	i, readLen, err := int32IDFromAZIDBinField(b, typeHint)
	if err == nil {
		*id = i
	}
	return readLen, err
}

func int32IDFromAZIDBinField(
	b []byte, typeHint azid.BinDataType,
) (id int32IDNum, readLen int, err error) {
	if typeHint != azid.BinDataTypeUnspecified && typeHint != azid.BinDataTypeInt32 {
		return int32IDNum(0), 0,
			errors.ArgMsg("typeHint", "unsupported")
	}
	i := binary.BigEndian.Uint32(b)
	return int32IDNum(i), 4, nil
}

type int32RefKey int32IDNum

var _ azid.BinMarshalable = int32RefKey(0)
var _ azid.BinFieldMarshalable = int32RefKey(0)

func int32RefKeyFromAZIDBinField(
	b []byte, typeHint azid.BinDataType,
) (refKey int32RefKey, readLen int, err error) {
	id, n, err := int32IDFromAZIDBinField(b, typeHint)
	if err != nil {
		return int32RefKey(0), n, err
	}
	return int32RefKey(id), n, nil
}

func (refKey int32RefKey) AZIDBin() []byte {
	b := make([]byte, 5)
	b[0] = azid.BinDataTypeInt32.Byte()
	binary.BigEndian.PutUint32(b[1:], uint32(refKey))
	return b
}

func int32RefKeyFromAZIDBin(b []byte) (refKey int32RefKey, readLen int, err error) {
	typ, err := azid.BinDataTypeFromByte(b[0])
	if err != nil {
		return int32RefKey(0), 0,
			errors.ArgWrap("", "type parsing", err)
	}
	if typ != azid.BinDataTypeInt32 {
		return int32RefKey(0), 0,
			errors.Arg("", errors.EntMsg("type", "unsupported"))
	}

	i, readLen, err := int32IDFromAZIDBinField(b[1:], typ)
	if err != nil {
		return int32RefKey(0), 0,
			errors.ArgWrap("", "id data parsing", err)
	}

	return int32RefKey(i), 1 + readLen, nil
}

func (refKey int32RefKey) AZIDBinField() ([]byte, azid.BinDataType) {
	return int32IDNum(refKey).AZIDBinField()
}

type adjunctIDNum int16

func adjunctIDFromAZIDBinField(
	b []byte, typeHint azid.BinDataType,
) (id adjunctIDNum, readLen int, err error) {
	if typeHint != azid.BinDataTypeUnspecified && typeHint != azid.BinDataTypeInt16 {
		return adjunctIDNum(0), 0,
			errors.ArgMsg("typeHint", "unsupported")
	}
	i := binary.BigEndian.Uint16(b)
	return adjunctIDNum(i), 2, nil
}

func (id adjunctIDNum) AZIDBinField() ([]byte, azid.BinDataType) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(id))
	return b, azid.BinDataTypeInt16
}

type adjunctRefKey struct {
	parent int32RefKey
	id     adjunctIDNum
}

const adjunctRefKeyFieldCount = 2

func adjunctRefKeyFromAZIDBinField(
	b []byte, typeHint azid.BinDataType,
) (refKey adjunctRefKey, readLen int, err error) {
	if typeHint != azid.BinDataTypeArray {
		return adjunctRefKey{}, 0,
			errors.Arg("", errors.EntMsg("type", "unsupported"))
	}

	arrayLen := int(b[0])
	if arrayLen != adjunctRefKeyFieldCount {
		return adjunctRefKey{}, 0,
			errors.Arg("", errors.EntMsg("field count", "mismatch"))
	}

	typeCursor := 1
	dataCursor := typeCursor + arrayLen

	parentType, err := azid.BinDataTypeFromByte(b[typeCursor])
	if err != nil {
		return adjunctRefKey{}, 0,
			errors.ArgWrap("", "parent type parsing", err)
	}
	typeCursor++
	parentRefKey, readLen, err := int32RefKeyFromAZIDBinField(b[dataCursor:], parentType)
	if err != nil {
		return adjunctRefKey{}, 0,
			errors.ArgWrap("", "parent data parsing", err)
	}
	dataCursor += readLen

	idType, err := azid.BinDataTypeFromByte(b[typeCursor])
	if err != nil {
		return adjunctRefKey{}, 0,
			errors.ArgWrap("", "id type parsing", err)
	}
	typeCursor++
	id, readLen, err := adjunctIDFromAZIDBinField(b[dataCursor:], idType)
	if err != nil {
		return adjunctRefKey{}, 0,
			errors.ArgWrap("", "id data parsing", err)
	}
	dataCursor += readLen

	return adjunctRefKey{parentRefKey, id}, dataCursor, nil
}

func adjunctRefKeyFromAZIDBin(
	b []byte,
) (refKey adjunctRefKey, readLen int, err error) {
	typ, err := azid.BinDataTypeFromByte(b[0])
	if err != nil {
		return adjunctRefKey{}, 0,
			errors.ArgWrap("", "type parsing", err)
	}
	if typ != azid.BinDataTypeArray {
		return adjunctRefKey{}, 0,
			errors.Arg("", errors.EntMsg("type", "unsupported"))
	}

	refKey, readLen, err = adjunctRefKeyFromAZIDBinField(b[1:], typ)
	return refKey, readLen + 1, err
}

func (refKey adjunctRefKey) AZIDBinField() ([]byte, azid.BinDataType) {
	var typesBytes []byte
	var dataBytes []byte
	var fieldBytes []byte
	var fieldType azid.BinDataType

	fieldBytes, fieldType = refKey.parent.AZIDBinField()
	typesBytes = append(typesBytes, fieldType.Byte())
	dataBytes = append(dataBytes, fieldBytes...)

	fieldBytes, fieldType = refKey.id.AZIDBinField()
	typesBytes = append(typesBytes, fieldType.Byte())
	dataBytes = append(dataBytes, fieldBytes...)

	var out = []byte{byte(len(typesBytes))}
	out = append(out, typesBytes...)
	out = append(out, dataBytes...)
	return out, azid.BinDataTypeArray
}

func (refKey adjunctRefKey) AZIDBin() []byte {
	data, typ := refKey.AZIDBinField()
	out := []byte{typ.Byte()}
	return append(out, data...)
}

func TestEncodeField(t *testing.T) {
	testCases := []struct {
		in      int32IDNum
		outData []byte
		outType azid.BinDataType
	}{
		{int32IDNum(0), []byte{0, 0, 0, 0}, azid.BinDataTypeInt32},
		{int32IDNum(1), []byte{0, 0, 0, 1}, azid.BinDataTypeInt32},
		{int32IDNum(1 << 24), []byte{1, 0, 0, 0}, azid.BinDataTypeInt32},
	}

	for _, testCase := range testCases {
		outData, outType := testCase.in.AZIDBinField()
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
		outData := testCase.in.AZIDBin()
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
		{adjunctRefKey{int32RefKey(0), adjunctIDNum(0)},
			[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 0, 0, 0}},
		{adjunctRefKey{int32RefKey(0), adjunctIDNum(1)},
			[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 0, 0, 1}},
		{adjunctRefKey{int32RefKey(1), adjunctIDNum(0)},
			[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 1, 0, 0}},
	}

	for _, testCase := range testCases {
		outData := testCase.in.AZIDBin()
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
			adjunctRefKey{int32RefKey(0), adjunctIDNum(0)}, 10, nil},
		{[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 0, 0, 1},
			adjunctRefKey{int32RefKey(0), adjunctIDNum(1)}, 10, nil},
		{[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 1, 0, 0},
			adjunctRefKey{int32RefKey(1), adjunctIDNum(0)}, 10, nil},
	}

	for _, testCase := range testCases {
		refKey, readLen, err := adjunctRefKeyFromAZIDBin(testCase.in)
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
