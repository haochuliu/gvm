package classfile

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlePc  uint16
	cacheType uint16
}

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (this *CodeAttribute) readInfo(reader *ClassReader) {
	this.maxStack = reader.readUint16()
	this.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	this.code = reader.readBytes(codeLength)
	this.exceptionTable = readExceptionTable(reader)
	this.attributes = readAttributes(reader, this.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlePc:  reader.readUint16(),
			cacheType: reader.readUint16(),
		}
	}
	return exceptionTable
}
