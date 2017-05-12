package data

import (
	"encoding/hex"
	"fmt"
)

type MacAddr []byte

func (m MacAddr) IsValid() bool {
	return len(m) == 6
}

func (m MacAddr) String() string {
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		m[0], m[1], m[2], m[3], m[4], m[5])
}

func (m MacAddr) HexString() string {
	return hex.EncodeToString(m)
}

func (m MacAddr) MarshalJSON() ([]byte, error) {
	return []byte(`"` + m.String() + `"`), nil
}
