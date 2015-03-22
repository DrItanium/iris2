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

// functions for the pointer type
package iris2

import (
	"bytes"
	"encoding/binary"
)

type Pointer interface {
	Address() Word
	ReadByte() byte
	ReadQuarterWord(binary.ByteOrder) (QuarterWord, error)
	ReadHalfWord(binary.ByteOrder) (HalfWord, error)
	ReadWord(binary.ByteOrder) (Word, error)
	ReadFloat(binary.ByteOrder) (Float, error)
	ReadDouble(binary.ByteOrder) (Double, error)
	WriteByte(byte)
	WriteQuarterWord(QuarterWord, binary.ByteOrder) error
	WriteHalfWord(HalfWord, binary.ByteOrder) error
	WriteWord(Word, binary.ByteOrder) error
	WriteFloat(Float, binary.ByteOrder) error
	WriteDouble(Double, binary.ByteOrder) error
	// Used for instruction parsing
	ByteReader() *bytes.Reader
}

type StandardPointer struct {
	raw             []byte
	StartingAddress Word
}

func NewStandardPointer(rawPtr []byte, address Word) *StandardPointer {
	var p StandardPointer
	p.raw = rawPtr
	p.StartingAddress = address
	return &p
}
func (this *StandardPointer) Address() Word {
	return this.StartingAddress
}
func (this *StandardPointer) ReadByte() byte {
	return this.raw[0]
}

func (this *StandardPointer) ReadQuarterWord(order binary.ByteOrder) (QuarterWord, error) {
	var qw QuarterWord
	buf := bytes.NewReader(this.raw)
	err := binary.Read(buf, order, &qw)
	if err != nil {
		qw = 0
	}
	return qw, err
}

func (this *StandardPointer) ReadHalfWord(order binary.ByteOrder) (HalfWord, error) {
	var hw HalfWord
	buf := bytes.NewReader(this.raw)
	err := binary.Read(buf, order, &hw)
	if err != nil {
		hw = 0
	}
	return hw, err
}

func (this *StandardPointer) ReadWord(order binary.ByteOrder) (Word, error) {
	var w Word
	buf := bytes.NewReader(this.raw)
	err := binary.Read(buf, order, &w)
	if err != nil {
		w = 0
	}
	return w, err
}

func (this *StandardPointer) ReadFloat(order binary.ByteOrder) (Float, error) {
	var f Float
	result, err := this.ReadHalfWord(order)
	if err == nil {
		f = result.Float()
	} else {
		f = 0.0
	}
	return f, err
}

func (this *StandardPointer) ReadDouble(order binary.ByteOrder) (Double, error) {
	var d Double
	result, err := this.ReadWord(order)
	if err == nil {
		d = result.Double()
	} else {
		d = 0.0
	}
	return d, err
}

func (this *StandardPointer) WriteByte(value byte) {
	this.raw[0] = value
}

func (this *StandardPointer) WriteQuarterWord(value QuarterWord, order binary.ByteOrder) error {
	order.PutUint16(this.raw, uint16(value))
	return nil
}
func (this *StandardPointer) WriteHalfWord(value HalfWord, order binary.ByteOrder) error {
	order.PutUint32(this.raw, uint32(value))
	return nil
}
func (this *StandardPointer) WriteWord(value Word, order binary.ByteOrder) error {
	order.PutUint64(this.raw, uint64(value))
	return nil
}

func (this *StandardPointer) WriteFloat(value Float, order binary.ByteOrder) error {
	order.PutUint32(this.raw, uint32(value.BinaryRepresentation()))
	return nil
}
func (this *StandardPointer) WriteDouble(value Double, order binary.ByteOrder) error {
	order.PutUint64(this.raw, uint64(value.BinaryRepresentation()))
	return nil
}

// Used for instruction parsing
func (this *StandardPointer) ByteReader() *bytes.Reader {
	return bytes.NewReader(this.raw)
}
