package stss

import (
	"encoding/binary"
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/fullbox"
)

const (
	STSS string = "stss"
)

type Box struct {
	fullbox.Box
	SampleCount  uint32
	SampleNumber []uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{
		*fullbox.New(i),
		0,
		make([]uint32, 0, 0),
	}
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s, samplecount=%d, samplenumber=%+v",
		b.Info(),
		b.SampleCount,
		b.SampleNumber,
	)
}

func (Box) Type() string {
	return STSS
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)
	b.SampleCount = sr.Uint32()
	b.SampleNumber = make([]uint32, b.SampleCount)

	for i := uint32(0); i < b.SampleCount; i++ {
		b.SampleNumber[i] = binary.BigEndian.Uint32(src[(8 + 4*i):(12 + 4*i)])
	}

	return box.FullRead(len(src))
}
