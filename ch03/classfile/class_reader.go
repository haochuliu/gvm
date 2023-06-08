package classfile

import "encoding/binary"

const (
	U8 = 1 << iota
	U16
	U32
	U64
)

type ClassReader struct {
	data []byte
}

func (this *ClassReader) readUint8() uint8 {

	val := this.data[0]
	this.data = this.data[U8:]
	return val
}

func (this *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(this.data)
	this.data = this.data[U16:]
	return val
}

func (this *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(this.data)
	this.data = this.data[U32:]
	return val
}

func (this *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(this.data)
	this.data = this.data[U64:]
	return val
}

func (this *ClassReader) readUint16s() []uint16 {
	length := this.readUint16()
	result := make([]uint16, length)
	for i := range result {
		result[i] = this.readUint16()
	}
	return result
}

func (this *ClassReader) readBytes(n uint32) []byte {
	bytes := this.data[:n]
	this.data = this.data[n:]
	return bytes
}
