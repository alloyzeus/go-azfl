package azer_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/alloyzeus/go-azfl/azfl/azer"
	"github.com/alloyzeus/go-azfl/azfl/errors"
)

type int32ID int32

var _ azer.BinFieldMarshalable = int32ID(0)

func (id int32ID) AZERBinField() ([]byte, azer.BinDataType) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(id))
	return b, azer.BinDataTypeInt32
}

func (id *int32ID) UnmarshalAZERBinField(b []byte, typeHint azer.BinDataType) (readLen int, err error) {
	i, readLen, err := int32IDFromAZERBinField(b, typeHint)
	if err == nil {
		*id = i
	}
	return readLen, err
}

func int32IDFromAZERBinField(
	b []byte, typeHint azer.BinDataType,
) (id int32ID, readLen int, err error) {
	if typeHint != azer.BinDataTypeUnspecified && typeHint != azer.BinDataTypeInt32 {
		return int32ID(0), 0,
			errors.ArgMsg("typeHint", "unsupported")
	}
	i := binary.BigEndian.Uint32(b)
	return int32ID(i), 4, nil
}

type int32RefKey int32ID

var _ azer.BinMarshalable = int32RefKey(0)
var _ azer.BinFieldMarshalable = int32RefKey(0)

func int32RefKeyFromAZERBinField(
	b []byte, typeHint azer.BinDataType,
) (refKey int32RefKey, readLen int, err error) {
	id, n, err := int32IDFromAZERBinField(b, typeHint)
	if err != nil {
		return int32RefKey(0), n, err
	}
	return int32RefKey(id), n, nil
}

func (refKey int32RefKey) AZERBin() []byte {
	b := make([]byte, 5)
	b[0] = azer.BinDataTypeInt32.Byte()
	binary.BigEndian.PutUint32(b[1:], uint32(refKey))
	return b
}

func int32RefKeyFromAZERBin(b []byte) (refKey int32RefKey, readLen int, err error) {
	typ, err := azer.BinDataTypeFromByte(b[0])
	if err != nil {
		return int32RefKey(0), 0,
			errors.ArgWrap("", "type parsing", err)
	}
	if typ != azer.BinDataTypeInt32 {
		return int32RefKey(0), 0,
			errors.Arg("", errors.EntMsg("type", "unsupported"))
	}

	i, readLen, err := int32IDFromAZERBinField(b[1:], typ)
	if err != nil {
		return int32RefKey(0), 0,
			errors.ArgWrap("", "id data parsing", err)
	}

	return int32RefKey(i), 1 + readLen, nil
}

func (refKey int32RefKey) AZERBinField() ([]byte, azer.BinDataType) {
	return int32ID(refKey).AZERBinField()
}

type adjunctID int16

func adjunctIDFromAZERBinField(
	b []byte, typeHint azer.BinDataType,
) (id adjunctID, readLen int, err error) {
	if typeHint != azer.BinDataTypeUnspecified && typeHint != azer.BinDataTypeInt16 {
		return adjunctID(0), 0,
			errors.ArgMsg("typeHint", "unsupported")
	}
	i := binary.BigEndian.Uint16(b)
	return adjunctID(i), 2, nil
}

func (id adjunctID) AZERBinField() ([]byte, azer.BinDataType) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(id))
	return b, azer.BinDataTypeInt16
}

type adjunctRefKey struct {
	parent int32RefKey
	id     adjunctID
}

const adjunctRefKeyFieldCount = 2

func adjunctRefKeyFromAZERBinField(
	b []byte, typeHint azer.BinDataType,
) (refKey adjunctRefKey, readLen int, err error) {
	if typeHint != azer.BinDataTypeArray {
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

	parentType, err := azer.BinDataTypeFromByte(b[typeCursor])
	if err != nil {
		return adjunctRefKey{}, 0,
			errors.ArgWrap("", "parent type parsing", err)
	}
	typeCursor++
	parentRefKey, readLen, err := int32RefKeyFromAZERBinField(b[dataCursor:], parentType)
	if err != nil {
		return adjunctRefKey{}, 0,
			errors.ArgWrap("", "parent data parsing", err)
	}
	dataCursor += readLen

	idType, err := azer.BinDataTypeFromByte(b[typeCursor])
	if err != nil {
		return adjunctRefKey{}, 0,
			errors.ArgWrap("", "id type parsing", err)
	}
	typeCursor++
	id, readLen, err := adjunctIDFromAZERBinField(b[dataCursor:], idType)
	if err != nil {
		return adjunctRefKey{}, 0,
			errors.ArgWrap("", "id data parsing", err)
	}
	dataCursor += readLen

	return adjunctRefKey{parentRefKey, id}, dataCursor, nil
}

func adjunctRefKeyFromAZERBin(
	b []byte,
) (refKey adjunctRefKey, readLen int, err error) {
	typ, err := azer.BinDataTypeFromByte(b[0])
	if err != nil {
		return adjunctRefKey{}, 0,
			errors.ArgWrap("", "type parsing", err)
	}
	if typ != azer.BinDataTypeArray {
		return adjunctRefKey{}, 0,
			errors.Arg("", errors.EntMsg("type", "unsupported"))
	}

	refKey, readLen, err = adjunctRefKeyFromAZERBinField(b[1:], typ)
	return refKey, readLen + 1, err
}

func (refKey adjunctRefKey) AZERBinField() ([]byte, azer.BinDataType) {
	var typesBytes []byte
	var dataBytes []byte
	var fieldBytes []byte
	var fieldType azer.BinDataType

	fieldBytes, fieldType = refKey.parent.AZERBinField()
	typesBytes = append(typesBytes, fieldType.Byte())
	dataBytes = append(dataBytes, fieldBytes...)

	fieldBytes, fieldType = refKey.id.AZERBinField()
	typesBytes = append(typesBytes, fieldType.Byte())
	dataBytes = append(dataBytes, fieldBytes...)

	var out = []byte{byte(len(typesBytes))}
	out = append(out, typesBytes...)
	out = append(out, dataBytes...)
	return out, azer.BinDataTypeArray
}

func (refKey adjunctRefKey) AZERBin() []byte {
	data, typ := refKey.AZERBinField()
	out := []byte{typ.Byte()}
	return append(out, data...)
}

func TestEncodeField(t *testing.T) {
	testCases := []struct {
		in      int32ID
		outData []byte
		outType azer.BinDataType
	}{
		{int32ID(0), []byte{0, 0, 0, 0}, azer.BinDataTypeInt32},
		{int32ID(1), []byte{0, 0, 0, 1}, azer.BinDataTypeInt32},
		{int32ID(1 << 24), []byte{1, 0, 0, 0}, azer.BinDataTypeInt32},
	}

	for _, testCase := range testCases {
		outData, outType := testCase.in.AZERBinField()
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
		outData := testCase.in.AZERBin()
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
		outData := testCase.in.AZERBin()
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
		refKey, readLen, err := adjunctRefKeyFromAZERBin(testCase.in)
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
