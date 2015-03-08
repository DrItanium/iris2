package iris2

import "math"

type Word uint64
type HalfWord uint32
type DoubleWord [2]Word
type QuadWord [4]Word
type QuarterWord uint16
type Register Word

func (this *Word) AsFloat64() float64 {
	return math.Float64frombits(uint64(*this))
}
func (this *Word) EncodeFloat64(value float64) {
	*this = Word(math.Float64bits(value))
}

func (this *HalfWord) AsFloat32() float32 {
	return math.Float32frombits(uint32(*this))
}

func (this *HalfWord) EncodeFloat32(value float32) {
	*this = HalfWord(math.Float32bits(value))
}

func (this *Word) IntegerHalves() []HalfWord {
	return []HalfWord{HalfWord(*this), HalfWord(*this >> 32)}
}
