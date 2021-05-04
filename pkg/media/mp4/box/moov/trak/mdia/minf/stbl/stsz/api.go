// Package stsz (Sample Size)
package stsz

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	STSZ string = "stsz"
)

type Box struct {
	base.Box
	SampleSize  uint32
	SampleCount uint32
	EntrySize   []uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, 0, 0, make([]uint32, 0)}
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
