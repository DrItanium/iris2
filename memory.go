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

// declaration of the iris2 memory layout
package iris2

import (
	"fmt"
)

// Memory is laid out as a series of spaces (which are slices). This makes it really easy to simulate NUMA designs where we have multiple discontiguous memory spaces
// The actual design of memory itself is abstracted through the use of the memory spaces and is up to the memory controller to actually manage things like address translation
// The simplest design will have a single memory space
type MemorySpace interface {
	MemorySize() Word
	PointerAt(address Word) (*Pointer, error)
	BoundedPointerAt(address, size Word) (*Pointer, error)
	Subspace(start Word) (MemorySpace, error)
	BoundedSubspace(start, end Word) (MemorySpace, error)
}

type SimpleMemorySpace []byte

// Returns slice for a given address within this space. The external address is not the spaceAddress in cases with multiple memory spaces
func (this SimpleMemorySpace) BoundedSubspace(start, end Word) (MemorySpace, error) {
	return SimpleMemorySpace(this[start:end]), nil
}
func (this SimpleMemorySpace) Subspace(start Word) (MemorySpace, error) {
	return SimpleMemorySpace(this[start:]), nil
}

func (this SimpleMemorySpace) MemorySize() Word {
	return Word(len(this))
}

func (this SimpleMemorySpace) PointerAt(address Word) (*Pointer, error) {
	if address >= Word(len(this)) {
		return nil, fmt.Errorf("Provided address %d is out of range!", address)
	} else {
		return NewPointer(this[address:], address), nil
	}
}

func (this SimpleMemorySpace) BoundedPointerAt(address, size Word) (*Pointer, error) {
	if address >= Word(len(this)) {
		return nil, fmt.Errorf("Provided address %d is out of range!", address)
	}

	// hmmm, I have to think if there are other kinds of failure states
	if size > Word(len(this)) || (size > (Word(len(this)) - address)) {
		return nil, fmt.Errorf("Provided bounding size %d is too big!", size)
	}

	return NewPointer(this[address:(address+size)], address), nil
}
