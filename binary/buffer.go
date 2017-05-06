package binary

type Buffer []byte

func NewBuffer(length uint) Buffer {
	return make(Buffer, length)
}

func (b Buffer) Write(offset uint, bytes []byte) {
	copy(b[offset:], bytes)
}

func (b Buffer) WriteUInt16BE(offset uint, v uint16) {
	b[offset] = byte((v & 0xff00) >> 8)
	b[offset+1] = byte(v & 0x00ff)
}

func (b Buffer) WriteUInt32BE(offset uint, v uint32) {
	b[offset] = byte((v & 0xff000000) >> 24)
	b[offset+1] = byte((v & 0x00ff0000) >> 16)
	b[offset+2] = byte((v & 0x0000ff00) >> 8)
	b[offset+3] = byte(v & 0x000000ff)
}
