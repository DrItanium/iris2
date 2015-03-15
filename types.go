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

// system data types
package iris2

import "math"

type Double float64
type Float float32
type Word uint64
type HalfWord uint32
type QuarterWord uint16

type ByteConvertible interface {
	Bytes() []byte
}

func (this Word) Double() Double {
	return Double(math.Float64frombits(uint64(this)))
}
func (this HalfWord) Float() Float {
	return Float(math.Float32frombits(uint32(this)))
}
func (this Float) BinaryRepresentation() HalfWord {
	return HalfWord(math.Float32bits(float32(this)))
}
func (this Double) BinaryRepresentation() Word {
	return Word(math.Float64bits(float64(this)))
}

func (this Word) Bytes() []byte {
	return []byte{byte(this), byte(this >> 8), byte(this >> 16), byte(this >> 24), byte(this >> 32), byte(this >> 40), byte(this >> 48), byte(this >> 56)}
}

func (this HalfWord) Bytes() []byte {
	return []byte{byte(this), byte(this >> 8), byte(this >> 16), byte(this >> 24)}
}

func (this Float) Bytes() []byte {
	return this.BinaryRepresentation().Bytes()
}

func (this Double) Bytes() []byte {
	return this.BinaryRepresentation().Bytes()
}
