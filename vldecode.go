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

// variable-length decoder
package iris2

import "fmt"

type VariableLengthInstruction struct {
	Contents []InstructionField
}

func (this *VariableLengthInstruction) Bytes() []byte {
	var output []byte
	for _, inst := range this.Contents {
		output = append(output, inst...)
	}
	return output
}

func (this *VariableLengthInstruction) Fields() []InstructionField {
	return this.Contents
}

func (this *VariableLengthInstruction) RawRepresentation() RawInstruction {
	return RawInstruction(this.Bytes())
}

// First style variable length decoder, iris1 style but with variable length constant fields
// That means that we will consume at most 28 bytes or four fields
type VariableLengthDecoder_Type0 struct{}

func (this *VariableLengthDecoder_Type0) Decode(ptr *Pointer) (Instruction, *Pointer, error) {
	// the first byte is used to describe the operation
	var vli VariableLengthInstruction
	reader := ptr.ByteReader()
	control := make(InstructionField, 1)
	vli.Contents = make([]InstructionField, 4)
	vli.Contents[0] = control
	count, err := reader.Read(control)
	if err != nil {
		return nil, ptr, err
	}
	if count == 0 {
		return nil, ptr, fmt.Errorf("Couldn't read from provided pointer to decode!")
	}
	return &vli, ptr, nil
}
