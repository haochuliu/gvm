package classfile

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (this *SourceFileAttribute) readInfo(reader *ClassReader) {
	this.sourceFileIndex = reader.readUint16()
}

func (this *SourceFileAttribute) FileName() string {
	return this.cp.getUtf8(this.sourceFileIndex)
}
