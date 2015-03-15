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

// generic instruction decoder
package iris2

type InstructionDecoder interface {
	Decode(ptr *Pointer) (*Instruction, *Pointer, error)
}
type Instruction interface {
	ByteConvertible
	RawRepresentation() RawInstruction
	Fields() []InstructionField
}
type RawInstruction []byte
type InstructionField []byte

func (this RawInstruction) Bytes() []byte {
	return []byte(this)
}

func (this RawInstruction) RawRepresentation() RawInstruction {
	return this
}

func (this RawInstruction) Fields() []InstructionField {
	contents := make([]InstructionField, len(this))
	for i := 0; i < len(this); i++ {
		contents[i] = InstructionField(this[i : i+1])
	}
	return contents
}
