package classfile

type MarkerAttribute struct{}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

func (this *MarkerAttribute) readInfo(reader *ClassReader) {}
