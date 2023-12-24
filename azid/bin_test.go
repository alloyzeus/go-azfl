package azid_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/alloyzeus/go-azfl/v2/azid"
	"github.com/alloyzeus/go-azfl/v2/errors"
)

type int32IDNum int32

var _ azid.BinFieldMarshalable = int32IDNum(0)

func (id int32IDNum) AZIDBinField() ([]byte, azid.BinDataType) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(id))
	return b, azid.BinDataTypeInt32
}

func (id *int32IDNum) UnmarshalAZIDBinField(b []byte, typeHint azid.BinDataType) (readLen int, err error) {
	i, readLen, err := int32IDNumFromAZIDBinField(b, typeHint)
	if err == nil {
		*id = i
	}
	return readLen, err
}

func int32IDNumFromAZIDBinField(
	b []byte, typeHint azid.BinDataType,
) (id int32IDNum, readLen int, err error) {
	if typeHint != azid.BinDataTypeUnspecified && typeHint != azid.BinDataTypeInt32 {
		return int32IDNum(0), 0,
			errors.ArgValueUnsupported("typeHint")
	}
	i := binary.BigEndian.Uint32(b)
	return int32IDNum(i), 4, nil
}

type int32ID int32IDNum

var _ azid.BinMarshalable = int32ID(0)
var _ azid.BinFieldMarshalable = int32ID(0)

func int32IDFromAZIDBinField(
	b []byte, typeHint azid.BinDataType,
) (id int32ID, readLen int, err error) {
	idNum, n, err := int32IDNumFromAZIDBinField(b, typeHint)
	if err != nil {
		return int32ID(0), n, err
	}
	return int32ID(idNum), n, nil
}

func (id int32ID) AZIDBin() []byte {
	b := make([]byte, 5)
	b[0] = azid.BinDataTypeInt32.Byte()
	binary.BigEndian.PutUint32(b[1:], uint32(id))
	return b
}

func int32IDFromAZIDBin(idBin []byte) (id int32ID, readLen int, err error) {
	typ, err := azid.BinDataTypeFromByte(idBin[0])
	if err != nil {
		return int32ID(0), 0,
			errors.Arg("idBin").Fieldset(errors.
				NamedValueMalformed("type").Wrap(err))
	}
	if typ != azid.BinDataTypeInt32 {
		return int32ID(0), 0,
			errors.Arg("idBin").Fieldset(errors.
				NamedValueUnsupported("type"))
	}

	i, readLen, err := int32IDNumFromAZIDBinField(idBin[1:], typ)
	if err != nil {
		return int32ID(0), 0,
			errors.Arg("idBin").Fieldset(errors.
				NamedValueMalformed("id data").Wrap(err))
	}

	return int32ID(i), 1 + readLen, nil
}

func (id int32ID) AZIDBinField() ([]byte, azid.BinDataType) {
	return int32IDNum(id).AZIDBinField()
}

type adjunctIDNum int16

func adjunctIDNumFromAZIDBinField(
	b []byte, typeHint azid.BinDataType,
) (id adjunctIDNum, readLen int, err error) {
	if typeHint != azid.BinDataTypeUnspecified && typeHint != azid.BinDataTypeInt16 {
		return adjunctIDNum(0), 0,
			errors.ArgValueUnsupported("typeHint")
	}
	i := binary.BigEndian.Uint16(b)
	return adjunctIDNum(i), 2, nil
}

func (id adjunctIDNum) AZIDBinField() ([]byte, azid.BinDataType) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(id))
	return b, azid.BinDataTypeInt16
}

type adjunctID struct {
	parent int32ID
	idNum  adjunctIDNum
}

const adjunctIDFieldCount = 2

func adjunctIDFromAZIDBinField(
	idBinField []byte, typeHint azid.BinDataType,
) (id adjunctID, readLen int, err error) {
	if typeHint != azid.BinDataTypeArray {
		return adjunctID{}, 0,
			errors.ArgValueUnsupported("typeHint")
	}

	arrayLen := int(idBinField[0])
	if arrayLen != adjunctIDFieldCount {
		return adjunctID{}, 0,
			errors.Arg("idBinField").Fieldset(errors.
				N("field count").Desc(errors.ErrValueMismatch))
	}

	typeCursor := 1
	dataCursor := typeCursor + arrayLen

	parentType, err := azid.BinDataTypeFromByte(idBinField[typeCursor])
	if err != nil {
		return adjunctID{}, 0,
			errors.Arg("idBinField").Fieldset(errors.
				NamedValueMalformed("parent type").Wrap(err))
	}
	typeCursor++
	parentID, readLen, err := int32IDFromAZIDBinField(idBinField[dataCursor:], parentType)
	if err != nil {
		return adjunctID{}, 0,
			errors.Arg("idBinField").Fieldset(errors.
				NamedValueMalformed("parent data").Wrap(err))
	}
	dataCursor += readLen

	idType, err := azid.BinDataTypeFromByte(idBinField[typeCursor])
	if err != nil {
		return adjunctID{}, 0,
			errors.Arg("idBinField").Fieldset(errors.
				NamedValueMalformed("id type").Wrap(err))
	}
	typeCursor++
	idNum, readLen, err := adjunctIDNumFromAZIDBinField(idBinField[dataCursor:], idType)
	if err != nil {
		return adjunctID{}, 0,
			errors.Arg("idBinField").Fieldset(errors.
				NamedValueMalformed("id data").Wrap(err))
	}
	dataCursor += readLen

	return adjunctID{parentID, idNum}, dataCursor, nil
}

func adjunctIDFromAZIDBin(
	idBin []byte,
) (id adjunctID, readLen int, err error) {
	typ, err := azid.BinDataTypeFromByte(idBin[0])
	if err != nil {
		return adjunctID{}, 0,
			errors.Arg("idBin").Fieldset(errors.
				NamedValueMalformed("type").Wrap(err))
	}
	if typ != azid.BinDataTypeArray {
		return adjunctID{}, 0,
			errors.Arg("idBin").Fieldset(errors.
				NamedValueUnsupported("type"))
	}

	id, readLen, err = adjunctIDFromAZIDBinField(idBin[1:], typ)
	return id, readLen + 1, err
}

func (id adjunctID) AZIDBinField() ([]byte, azid.BinDataType) {
	var typesBytes []byte
	var dataBytes []byte
	var fieldBytes []byte
	var fieldType azid.BinDataType

	fieldBytes, fieldType = id.parent.AZIDBinField()
	typesBytes = append(typesBytes, fieldType.Byte())
	dataBytes = append(dataBytes, fieldBytes...)

	fieldBytes, fieldType = id.idNum.AZIDBinField()
	typesBytes = append(typesBytes, fieldType.Byte())
	dataBytes = append(dataBytes, fieldBytes...)

	var out = []byte{byte(len(typesBytes))}
	out = append(out, typesBytes...)
	out = append(out, dataBytes...)
	return out, azid.BinDataTypeArray
}

func (id adjunctID) AZIDBin() []byte {
	data, typ := id.AZIDBinField()
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

func TestEncodeID(t *testing.T) {
	testCases := []struct {
		in      int32ID
		outData []byte
	}{
		{int32ID(0), []byte{0x13, 0, 0, 0, 0}},
		{int32ID(1), []byte{0x13, 0, 0, 0, 1}},
		{int32ID(1 << 24), []byte{0x13, 1, 0, 0, 0}},
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
		in      adjunctID
		outData []byte
	}{
		{adjunctID{int32ID(0), adjunctIDNum(0)},
			[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 0, 0, 0}},
		{adjunctID{int32ID(0), adjunctIDNum(1)},
			[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 0, 0, 1}},
		{adjunctID{int32ID(1), adjunctIDNum(0)},
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
		out     adjunctID
		readLen int
		err     error
	}{
		{[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 0, 0, 0},
			adjunctID{int32ID(0), adjunctIDNum(0)}, 10, nil},
		{[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 0, 0, 1},
			adjunctID{int32ID(0), adjunctIDNum(1)}, 10, nil},
		{[]byte{0x40, 0x2, 0x13, 0x12, 0, 0, 0, 1, 0, 0},
			adjunctID{int32ID(1), adjunctIDNum(0)}, 10, nil},
	}

	for _, testCase := range testCases {
		id, readLen, err := adjunctIDFromAZIDBin(testCase.in)
		if err != testCase.err {
			t.Errorf("Expected: %#v, actual: %#v", testCase.err, err)
		}
		if readLen != testCase.readLen || readLen != len(testCase.in) {
			t.Errorf("Expected: %#v, actual: %#v", testCase.readLen, readLen)
		}
		if id.parent != testCase.out.parent || id.idNum != testCase.out.idNum {
			t.Errorf("Expected: %#v, actual: %#v", testCase.out, id)
		}
	}
}
