package classfile

type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (this *ExceptionsAttribute) readInfo(reader *ClassReader) {
	this.exceptionIndexTable = reader.readUint16s()
}

func (this *ExceptionsAttribute) ExceptionTable() []uint16 {
	return this.exceptionIndexTable
}
