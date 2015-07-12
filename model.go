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

// iris core designs
package iris2

type FirstGenCore struct {
	Registers                 [256]Word
	OnDieMemory               SimpleMemorySpace // memory + pointers
	Code, Data, IO, Microcode MemorySpace
}

func Gigabyte(count Word) Word {
	return 1024 * Megabyte(count)
}
func Megabyte(count Word) Word {
	return 1024 * Kilobyte(count)
}
func Kilobyte(count Word) Word {
	return 1024 * count
}
func NewFirstGenCore(byteCount Word) *FirstGenCore {
	var core FirstGenCore
	for i := 0; i < 256; i++ {
		core.Registers[i] = 0
	}
	core.OnDieMemory = make([]byte, byteCount)
	core.Code = core.OnDieMemory[0 : byteCount/4]
	core.Data = core.OnDieMemory[byteCount/4 : byteCount/2]
	core.Stack = core.OnDieMemory[byteCount/2 : (byteCount*3)/4]
	core.IO = core
	return &core
}

type KeyboardController struct {
	CurrentData rune
}
