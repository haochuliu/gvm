package classfile

type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo
}

func (this *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	this.classIndex = reader.readUint16()
	this.nameAndTypeIndex = reader.readUint16()
}

func (this *ConstantMemberRefInfo) ClassName() string {
	return this.cp.getClassName(this.classIndex)
}

func (this *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return this.cp.getNameAndType(this.nameAndTypeIndex)
}
