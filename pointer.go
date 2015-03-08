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

type Pointer struct {
	raw     []byte
	Address Word
}

func NewPointer(rawPtr []byte, address Word) *Pointer {
	var p Pointer
	p.raw = rawPtr
	p.Address = address
	return &p
}
func (this *Pointer) ReadByte() byte {
	return this.raw[0]
}

func (this *Pointer) ReadQuarterWord(order binary.ByteOrder) (QuarterWord, error) {
	var qw QuarterWord
	buf := bytes.NewReader(this.raw)
	err := binary.Read(buf, order, &qw)
	if err != nil {
		qw = 0
	}
	return qw, err
}

func (this *Pointer) ReadHalfWord(order binary.ByteOrder) (HalfWord, error) {
	var hw HalfWord
	buf := bytes.NewReader(this.raw)
	err := binary.Read(buf, order, &hw)
	if err != nil {
		hw = 0
	}
	return hw, err
}

func (this *Pointer) ReadWord(order binary.ByteOrder) (Word, error) {
	var w Word
	buf := bytes.NewReader(this.raw)
	err := binary.Read(buf, order, &w)
	if err != nil {
		w = 0
	}
	return w, err
}

func (this *Pointer) ReadFloat(order binary.ByteOrder) (Float, error) {
	var f Float
	result, err := this.ReadHalfWord(order)
	if err == nil {
		f = result.Float()
	} else {
		f = 0.0
	}
	return f, err
}

func (this *Pointer) ReadDouble(order binary.ByteOrder) (Double, error) {
	var d Double
	result, err := this.ReadWord(order)
	if err == nil {
		d = result.Double()
	} else {
		d = 0.0
	}
	return d, err
}

func (this *Pointer) WriteByte(value byte) {
	this.raw[0] = value
}

func (this *Pointer) WriteQuarterWord(value QuarterWord, order binary.ByteOrder) error {
	buf := bytes.NewBuffer(this.raw)
	return binary.Write(buf, order, value)
}
func (this *Pointer) WriteHalfWord(value HalfWord, order binary.ByteOrder) error {
	buf := bytes.NewBuffer(this.raw)
	return binary.Write(buf, order, value)
}
func (this *Pointer) WriteWord(value Word, order binary.ByteOrder) error {
	buf := bytes.NewBuffer(this.raw)
	return binary.Write(buf, order, value)
}

func (this *Pointer) WriteFloat(value Float, order binary.ByteOrder) error {
	buf := bytes.NewBuffer(this.raw)
	return binary.Write(buf, order, value)
}
func (this *Pointer) WriteDouble(value Double, order binary.ByteOrder) error {
	buf := bytes.NewBuffer(this.raw)
	return binary.Write(buf, order, value)
}

// Used for instruction parsing
func (this *Pointer) ByteReader() *bytes.Reader {
	return bytes.NewReader(this.raw)
}
