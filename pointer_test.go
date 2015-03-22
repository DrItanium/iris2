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

import "testing"
import "encoding/binary"

func standardMemorySpace() SimpleMemorySpace {
	sms := make(SimpleMemorySpace, 128)
	for i := 0; i < 128; i++ {
		sms[i] = byte(i)
	}
	return sms
}
func defaultPointer(sms SimpleMemorySpace) (Pointer, error) {
	return sms.PointerAt(3)
}
func defaultSetup() (Pointer, error) {
	return defaultPointer(standardMemorySpace())
}
func TestReadByte(t *testing.T) {
	ptr, err := defaultSetup()
	if err != nil {
		t.Error(err)
	}
	brep := ptr.ReadByte()
	if brep != 3 {
		t.Errorf("Reading a byte returned: %d, not 3", brep)
	}
}

func TestReadQuarterWord(t *testing.T) {
	ptr, err := defaultSetup()
	if err != nil {
		t.Error(err)
	}
	qwrep, err1 := ptr.ReadQuarterWord(binary.LittleEndian)
	if err1 != nil {
		t.Error(err1)
	} else if qwrep != 0x0403 {
		t.Errorf("reading a quarter word returned: %X, not 0x0403", qwrep)
	}
}

func TestReadHalfWord(t *testing.T) {
	ptr, err := defaultSetup()
	if err != nil {
		t.Error(err)
	}

	hwrep, err2 := ptr.ReadHalfWord(binary.LittleEndian)
	if err2 != nil {
		t.Error(err2)
	} else if hwrep != 0x06050403 {
		t.Errorf("reading a half word returned: %X, not 0x06050403", hwrep)
	}
}

func TestReadWord(t *testing.T) {
	ptr, err := defaultSetup()
	if err != nil {
		t.Error(err)
	}

	wrep, err2 := ptr.ReadWord(binary.LittleEndian)
	if err2 != nil {
		t.Error(err2)
	} else if wrep != 0x0A09080706050403 {
		t.Errorf("reading a word returned: %X, not 0x0a09080706050403", wrep)
	}
}

func TestWriteByte(t *testing.T) {
	ptr, err := defaultSetup()
	if err != nil {
		t.Error(err)
	}
	ptr.WriteByte(32)
	if value := ptr.ReadByte(); value != 32 {
		t.Errorf("Writing 32 to address 3 was not successful, got %d instead", value)
	}
}

func TestWriteQuarterWord(t *testing.T) {
	ptr, err := defaultSetup()
	if err != nil {
		t.Error(err)
	}
	ptr.WriteQuarterWord(32, binary.LittleEndian)
	value, err0 := ptr.ReadQuarterWord(binary.LittleEndian)
	if err0 != nil {
		t.Error(err)
	} else if value != 32 {
		t.Errorf("Writing quarter word (32) to address 3 was not successful, got %d instead", value)
	}
}

func TestWriteHalfWord(t *testing.T) {
	sms := make(SimpleMemorySpace, 128)
	for i := 0; i < 128; i++ {
		sms[i] = byte(i)
	}
	ptr, err := sms.PointerAt(3)
	if err != nil {
		t.Error(err)
	}
	ptr.WriteHalfWord(2000000000, binary.LittleEndian)
	t.Logf("%v", sms)
	value, err0 := ptr.ReadHalfWord(binary.LittleEndian)
	if err0 != nil {
		t.Error(err)
	} else if value != 2000000000 {
		t.Logf("%v", sms)
		t.Errorf("Writing half word (2000000000) to address 3 was not successful, got %d instead", value)
	}
}

func TestWriteWord(t *testing.T) {
	sms := standardMemorySpace()
	ptr, err := sms.PointerAt(3)
	if err != nil {
		t.Error(err)
	}
	ptr.WriteWord(32, binary.LittleEndian)
	value, err0 := ptr.ReadWord(binary.LittleEndian)
	if err0 != nil {
		t.Error(err)
	} else if value != 32 {
		t.Logf("%v", sms)
		t.Errorf("Writing word (32) to address 3 was not successful, got %d instead", value)
	}
}
