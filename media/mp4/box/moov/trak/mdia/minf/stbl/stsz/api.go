// Package stsz (Sample Size)
package stsz

import (
	"encoding/binary"
	"fmt"
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	STSZ string = "stsz"
)

type Box struct {
	base.Box
	Version           byte
	Flags             uint32
	SampleUniformSize uint32
	SampleNumber      uint32
	SampleSize        []uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		0,
		make([]uint32, 0),
	}
}

func (Box) Type() string {
	return STSZ
}

func (b *Box) String() string {
	return fmt.Sprintf(
		"%s, version=%d, flag=s%d, uniformsize=%d, samples=%d, samplesize=%+v",
		b.Info(),
		b.Version,
		b.Flags,
		b.SampleUniformSize,
		b.SampleNumber,
		b.SampleSize,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	versionAndFlags := sr.Uint32()

	b.Version = byte(versionAndFlags >> 24)
	b.Flags = versionAndFlags & box.FlagsMask
	b.SampleUniformSize = sr.Uint32()
	b.SampleNumber = sr.Uint32()
	b.SampleSize = make([]uint32, b.SampleNumber)

	if sr.Length() > 12 {
		for i := 0; i < int(b.SampleNumber); i++ {
			b.SampleSize = append(b.SampleSize, binary.BigEndian.Uint32(src[(12+4*i):(16+4*i)]))
		}
	}

	return box.FullRead(len(src))
}
