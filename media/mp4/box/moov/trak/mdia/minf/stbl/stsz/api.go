// Package stsz (Sample Size)
package stsz

import (
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	STSZ string = "stsz"
)

type Box struct {
	base2.Box
	SampleSize  uint32
	SampleCount uint32
	EntrySize   []uint32
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, 0, 0, make([]uint32, 0)}
}

func (Box) Type() string {
	return STSZ
}

func (b *Box) String() string {
	return b.Info().String()
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
