package classfile

const (
	_ = iota
	CONSTANT_Utf8
	_
	CONSTANT_Integer
	CONSTANT_Float
	CONSTANT_Long
	CONSTANT_Double
	CONSTANT_Class
	CONSTANT_String
	CONSTANT_Fieldref
	CONSTANT_Methodref
	CONSTANT_InterfaceMethodref
	CONSTANT_NameAndType
	_
	_
	CONSTANT_MethodHandle
	CONSTANT_MethodType
	_
	CONSTANT_InvokeDynamic
)

/*
	cp_info {
	    u1 tag;
	    u1 info[];
	}
*/
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

// todo ugly code
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
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
