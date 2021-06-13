package stss

import (
	"encoding/binary"
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	STSS string = "stss"
)

type Box struct {
	base.Box
	Version      byte
	Flags        uint32
	SampleCount  uint32
	SampleNumber []uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		make([]uint32, 0, 0),
	}
}

func (b Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (Box) Type() string {
	return STSS
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	versionAndFlags := sr.Uint32()
	b.Version = byte(versionAndFlags >> 24)
	b.Flags = versionAndFlags & box.FlagsMask
	b.SampleCount = sr.Uint32()
	b.SampleNumber = make([]uint32, b.SampleCount)

	for i := 0; i < len(b.SampleNumber); i++ {
		b.SampleNumber = append(b.SampleNumber, binary.BigEndian.Uint32(src[(8+4*i):(12+4*i)]))
	}

	return box.FullRead(len(src))
}
