package rtda

import "math"

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make(LocalVars, maxLocals)
	}
	return nil
}

func (this LocalVars) SetInt(index uint, val int32) {
	this[index].num = val
}

func (this LocalVars) GetInt(index uint) int32 {
	return this[index].num
}

func (this LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	this[index].num = int32(bits)
}

func (this LocalVars) GetFloat(index uint) float32 {
	bits := uint32(this[index].num)
	return math.Float32frombits(bits)
}

func (this LocalVars) SetLong(index uint, val int64) {
	this[index].num = int32(val)
	this[index+1].num = int32(val >> 32)
}

func (this LocalVars) GetLong(index uint) int64 {
	low := uint32(this[index].num)
	high := uint32(this[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (this LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	this.SetLong(index, int64(bits))
}

func (this LocalVars) GetDouble(index uint) float64 {
	bits := uint64(this.GetLong(index))
	return math.Float64frombits(bits)
}

func (this LocalVars) SetRef(index uint, ref *Object) {
	this[index].ref = ref
}

func (this LocalVars) GetRef(index uint) *Object {
	return this[index].ref
}
