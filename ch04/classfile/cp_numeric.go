package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

func (this *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	this.val = int32(reader.readUint32())
}

type ConstantFloatInfo struct {
	val float32
}

func (this *ConstantFloatInfo) readInfo(reader *ClassReader) {
	this.val = math.Float32frombits(reader.readUint32())
}

type ConstantLongInfo struct {
	val int64
}

func (this *ConstantLongInfo) readInfo(reader *ClassReader) {
	this.val = int64(reader.readUint64())
}

type ConstantDoubleInfo struct {
	val float64
}

func (this *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	this.val = math.Float64frombits(reader.readUint64())
}
