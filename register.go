// Copyright (c) 2015 Joshua Scoggins
//
// This software is provided 'as-is', without any express or implied
// warranty. In no event will the authors be held liable for any damages
// arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it
// freely, subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented; you must not
//    claim that you wrote the original software. If you use this software
//    in a product, an acknowledgement in the product documentation would be
//    appreciated but is not required.
// 2. Altered source versions must be plainly marked as such, and must not be
//    misrepresented as being the original software.
// 3. This notice may not be removed or altered from any source distribution.
//
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
