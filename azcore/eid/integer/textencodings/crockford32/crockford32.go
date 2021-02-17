package crockford32

import (
	icrock32 "github.com/richardlehane/crock32"
)

func EncodeInt64(i int64) string {
	return icrock32.Encode(uint64(i))
}

func DecodeAsInt64(s string) (int64, error) {
	u, err := icrock32.Decode(s)
	return int64(u), err
}
