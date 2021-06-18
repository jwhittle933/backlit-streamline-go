package bits

import (
	"bytes"
	"fmt"
	"io"
)

const (
	Zero Bit = 0
	One  Bit = 1
)

type Bit byte
type Bits []Bit

type Reader struct {
	r      io.ByteReader
	src    []byte
	byte   byte
	offset uint8
}

func FromBool(b bool) Bit {
	if b {
		return One
	}

	return Zero
}

func (b Bit) Bool() bool {
	if b == 0 {
		return false
	}

	return true
}

func (b Bit) String() string {
	return fmt.Sprintf("%d", b)
}

func (b Bit) Byte() byte {
	return byte(b)
}

func NewReader(src []byte) *Reader {
	return &Reader{bytes.NewBuffer(src), make([]byte, 0), 0, 0}
}

func (r *Reader) ReadBit() (Bit, error) {
	if r.offset == 8 {
		r.offset = 0
	}

	if r.offset == 0 {
		var err error
		if r.byte, err = r.r.ReadByte(); err != nil {
			return Zero, err
		}

		r.captureByte()
	}

	bit := r.byte & (0x80 >> r.offset)
	r.offset++
	return FromBool(bit == 1), nil
}

func (r *Reader) ReadBits(n int64) (Bits, error) {
	bits := make([]Bit, n)

	for i := n - 1; i >= 0; i-- {
		bit, err := r.ReadBit()
		if err != nil {
			return bits, err
		}

		bits[i] = bit
	}

	return bits, nil
}

func (r *Reader) Bytes() []byte {
	return r.src
}

func (r *Reader) captureByte() {
	r.src = append(r.src, r.byte)
}
