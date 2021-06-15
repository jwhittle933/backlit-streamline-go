// Package saiz (Sample Auxiliary Information Sizes)
package saiz

import (
	"fmt"
	"github.com/jwhittle933/streamline/bits/slicereader"

	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/fullbox"
)

const (
	SAIZ string = "saiz"
)

type Box struct {
	fullbox.Box
	AuxInfoType           [4]byte
	AuxInfoTypeParameter  uint32
	DefaultSampleInfoSize byte
	SampleCount           uint32
	SampleInfo            []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{
		*fullbox.New(i),
		[4]byte{},
		0,
		0,
		0,
		make([]byte, 0, 0),
	}
}

func (Box) Type() string {
	return SAIZ
}

func (b Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)
	if b.Flags&0x01 != 0 {
		copy(b.AuxInfoType[:], sr.Slice(4))
		b.AuxInfoTypeParameter = sr.Uint32()
	}

	b.DefaultSampleInfoSize = sr.Uint8()
	b.SampleCount = sr.Uint32()
	b.SampleInfo = make([]byte, b.SampleCount)
	if b.DefaultSampleInfoSize == 0 {
		for i := 0; i < len(b.SampleInfo); i++ {
			b.SampleInfo[i] = sr.Uint8()
		}
	}

	return box.FullRead(len(src))
}
