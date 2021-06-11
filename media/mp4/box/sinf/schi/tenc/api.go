// Package tenc (Track Encryption)
package tenc

import (
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	TENC string = "tenc"
)

type Box struct {
	base2.Box
	_reserved              uint8
	DefaultCryptByteBlock  uint8
	DefaultSkipByteBlock   uint8
	DefaultIsProtected     uint8
	DefaultPerSampleIVSize uint8
	DefaultKID             [16]byte
	DefaultConstantIVSize  uint8
	DefaultConstantIV      []byte
	raw                    []byte
}

func New(i *box2.Info) box2.Boxed {
	return &Box{
		base2.Box{BoxInfo: i},
		0,
		0,
		0,
		0,
		0,
		[16]byte{},
		0,
		make([]byte, 0),
		make([]byte, 0),
	}
}

func (Box) Type() string {
	return TENC
}

func (b *Box) String() string {
	return b.Info().String()
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}