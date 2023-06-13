package classfile

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

func (this *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	this.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range this.lineNumberTable {
		this.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}

}
