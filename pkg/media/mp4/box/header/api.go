package header

import (
	"fmt"
)

const (
	MinSize int = 4
	MaxSize int = 8
)

var (
	headerSizeTooSmall = fmt.Errorf("source contains no header")
)

type Sizer interface {
	Size() int
	Raw() []byte
}

func New(src []byte) (Sizer, error) {
	// size is either [4]byte or [8]byte
	// if size is 0x00000000, the box continues until EOF
	if len(src) < MinSize {
		return nil, headerSizeTooSmall
	}

	if src[4] == 0x01 {
		return Byte8([8]byte{
			src[0], src[1], src[2], src[3],
			src[4], src[5], src[6], src[7],
		}), nil
	}

	return Byte4([4]byte{
		src[0], src[1], src[2], src[3],
	}), nil
}

type Byte4 [4]byte
type Byte8 [8]byte

func (b8 Byte8) Size() int {
	return 8
}

func (b8 Byte8) Raw() []byte {
	return b8[:]
}

func (b4 Byte4) Size() int {
	return 4
}

func (b4 Byte4) Raw() []byte {
	return b4[:]
}