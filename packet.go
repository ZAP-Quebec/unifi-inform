package inform

import (
	"errors"
	"github.com/ZAP-Quebec/unifi-inform/binary"
	messages "github.com/ZAP-Quebec/unifi-inform/data"
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

type Packet struct {
	ap    messages.MacAddr
	flags uint16
	key   messages.Key
	Msg   messages.Message
}

func NewPacket(ap messages.MacAddr, msg messages.Message, k messages.Key) *Packet {
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

	if p.IsZLib() {
		msg, err = CompressZLib(msg)
		if err != nil {
			return
		}
	} else if p.IsSnappy() {
		msg, err = CompressSnappy(msg)
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

	return b, nil
}

func (p *Packet) Unmarshal(data []byte, keyFetcher func(messages.MacAddr) (messages.Key, error)) (err error) {

	b := binary.Buffer(data)

	if len(data) < 40 {
		return errors.New("Invalid packet length.")
	}
	dataLength := uint(b.ReadUInt32BE(36))
	if uint(len(data)) < dataLength+40 {
		return errors.New("Invalid packet length.")
	}
	if b.ReadUInt32BE(0) != MAGIC_NUMBER {
		return errors.New("Invalid magic number at start of packet.")
	}
	if b.ReadUInt32BE(4) != INFORM_VERSION {
		return errors.New("Unkwown inform version.")
	}
	if b.ReadUInt32BE(32) != DATA_VERSION {
		return errors.New("Unkwown data version.")
	}

	p.ap = b.Read(8, 14)

	p.flags = b.ReadUInt16BE(14)

	msg := b.Read(40, 40+dataLength)
	if p.IsEncrypted() {
		iv := IV(b.Read(16, 32))
		p.key, err = keyFetcher(p.ap)
		if err != nil {
			return err
		}
		msg, err = Decrypt(iv, p.key, msg)
		if err != nil {
			return err
		}
	}

	if p.IsZLib() {
		msg, err = DecompressZLib(msg)
		if err != nil {
			return err
		}
	} else if p.IsSnappy() {
		msg, err = DecompressSnappy(msg)
		if err != nil {
			return err
		}
	}

	p.Msg, err = messages.Unmarshal(msg)
	return err
}
