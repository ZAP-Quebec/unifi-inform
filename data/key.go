package data

import (
	"bytes"
	"encoding/hex"
)

const KEY_SIZE = 16

type Key []byte

func (k Key) IsValid() bool {
	return len(k) == KEY_SIZE
}

func (k Key) IsDefault() bool {
	return bytes.Equal(k, DEFAULT_KEY)
}

func (k Key) String() string {
	return hex.EncodeToString(k)
}

var (
	DEFAULT_KEY = Key([]byte{
		0xba, 0x86, 0xf2, 0xbb,
		0xe1, 0x07, 0xc7, 0xc5,
		0x7e, 0xb5, 0xf2, 0x69,
		0x07, 0x75, 0xc7, 0x12,
	})
)
