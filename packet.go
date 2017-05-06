package inform

import (
	"errors"
	"fmt"
	"github.com/ZAP-Quebec/unifi-inform/binary"
)

const (
	MAGIC_NUMBER   uint32 = 1414414933
	INFORM_VERSION uint32 = 0
	DATA_VERSION   uint32 = 1

	ENCRYPT_FLAG uint16 = 1
	ZLIB_FLAG    uint16 = 2
	SNAPPY_FLAG  uint16 = 4
)

var (
	NIL_IV IV = []byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}
)

type MAC []byte

func (m MAC) IsValid() bool {
	return len(m) == 6
}

type Packet struct {
	ap    []byte
	flags uint16
	key   Key
	Msg   Message
}

func NewPacket(ap []byte, msg Message, k Key) *Packet {
	flags := SNAPPY_FLAG
	if k != nil {
		flags = flags | ENCRYPT_FLAG
	}
	return &Packet{
		ap:    ap,
		key:   k,
		Msg:   msg,
		flags: flags,
	}
}

func (p Packet) IsEncrypted() bool {
	return p.flags&ENCRYPT_FLAG == ENCRYPT_FLAG
}

func (p Packet) IsZLib() bool {
	return p.flags&ZLIB_FLAG == ZLIB_FLAG
}

func (p Packet) IsSnappy() bool {
	return p.flags&SNAPPY_FLAG == SNAPPY_FLAG
}

func (p Packet) Marshal() (result []byte, err error) {
	msg := p.Msg.Marshal()

	var l int
	if p.IsZLib() {
		msg, err = CompressZLib(msg)
		if err != nil {
			return
		}
	} else if p.IsSnappy() {
		l = len(msg)
		msg, err = CompressSnappy(msg)
		fmt.Printf("Snappy %d => %d \n", l, len(msg))
		if err != nil {
			return
		}
	}

	iv := NIL_IV
	if p.IsEncrypted() {
		iv, err = GenerateIV()
		if err != nil {
			return
		}
		msg, err = Encrypt(iv, p.key, msg)
		fmt.Printf("iv: %s key: %s \n", Key(iv).String(), p.key.String())
		if err != nil {
			return
		}
	}

	if len(p.ap) != 6 {
		return nil, errors.New("Invalid length of MAC address")
	}

	b := binary.NewBuffer(uint(40 + len(msg)))
	b.WriteUInt32BE(0, MAGIC_NUMBER)
	b.WriteUInt32BE(4, INFORM_VERSION)
	b.Write(8, p.ap)
	b.WriteUInt16BE(14, p.flags)
	b.Write(16, iv)
	b.WriteUInt32BE(32, DATA_VERSION)
	b.WriteUInt32BE(36, uint32(len(msg)))
	b.Write(40, msg)
	fmt.Printf("Msg: f:%d l:%d \n", p.flags, len(msg))

	return b, nil
}

func (p *Packet) Unmarshal(data []byte) error {
	return nil
}
