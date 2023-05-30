package classfile

type MemberInfo struct {
	cp               ContantPool
	accessFlag       uint16
	nameIndex        uint16
	descriptionIndex uint16
	attributes       []AttributeInfo
}

func readMembers(reader *ClassReader, cp ContantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ContantPool) *MemberInfo {
	return &MemberInfo{
		cp:               cp,
		accessFlag:       reader.readUint16(),
		nameIndex:        reader.readUint16(),
		descriptionIndex: reader.readUint16(),
		attributes:       readAttributes(reader, cp),
	}
}


