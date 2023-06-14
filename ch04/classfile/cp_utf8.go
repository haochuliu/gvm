package classfile

type ConstantUtf8Info struct {
	str string
}

func (this *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	this.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
