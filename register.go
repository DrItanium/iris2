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

type Register Word

func (this Register) UpperHalf() HalfWord {
	return HalfWord(this >> 32)
}
func (this Register) LowerHalf() HalfWord {
	return HalfWord(this)
}

func (this Register) Quarters() []QuarterWord {
	return []QuarterWord{
		QuarterWord(this),
		QuarterWord(this >> 16),
		QuarterWord(this >> 32),
		QuarterWord(this >> 48),
	}
}

func (this Register) Halves() []HalfWord {
	return []HalfWord{
		HalfWord(this),
		HalfWord(this >> 32),
	}
}

func (this Register) Bytes() []byte {
	return Word(this).Bytes()
}

func (this Register) Double() Double {
	return Word(this).Double()
}

func (this Register) Floats() []Float {
	return []Float{
		this.LowerHalf().Float(),
		this.UpperHalf().Float(),
	}
}
