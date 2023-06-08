package classfile

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

type ConstantPool []ConstantInfo

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{}
	case CONSTANT_Class:
		return &ConstantClassInfo{}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang,ClassFormatError: constant pool tag!")
	}
}

func readConstantPool(read *ClassReader) ConstantPool {
	cpCount := int(read.readUint16())
	cp := make(ConstantPool, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDubboInfo:
			i++
		}
	}

	return cp
}

func (this ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := this[index]; cpInfo != nil {
		return cpInfo
	}

	panic("Invalid constant pool index!")
}

func (this ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := this.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := this.getUtf8(ntInfo.nameIndex)
	_type := this.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (this ConstantPool) getClassName(index uint16) string {
	classInfo := this.getConstantInfo(index).(*ConstantClassInfo)
	return classInfo.getUtf8(classInfo.nameIndex)
}

func (this ConstantPool) getUtf8(index uint16) string {
	utf8Info := this.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
