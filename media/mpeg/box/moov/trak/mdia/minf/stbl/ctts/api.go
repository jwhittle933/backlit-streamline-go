package ctts

import (
	"encoding/binary"
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	CTTS string = "ctts"
)

type Box struct {
	base.Box
	Version      byte
	Flags        uint32
	SampleCount  []uint32
	SampleOffset []int32
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		make([]uint32, 0),
		make([]int32, 0),
	}
}

type Sample struct {
	SampleCount    uint32
	SampleOffsetV0 uint32
	SampleOffsetV1 int32
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s, version=%d, flags=%d, samplecount=%+v, sampleoffsets=%+v",
		b.Info(),
		b.Version,
		b.Flags,
		b.SampleCount,
		b.SampleOffset,
	)
}

func (Box) Type() string {
	return CTTS
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	versionAndFlags := sr.Uint32()
	b.Version = byte(versionAndFlags >> 24)
	b.Flags = versionAndFlags & box.FlagsMask

	for i, count := 0, sr.Uint32(); i < int(count); i++ {
		sCount := binary.BigEndian.Uint32(src[(8 + 8*i):(12 + 8*i)])
		sOffset := binary.BigEndian.Uint32(src[(12 + 8*i):(16 + 8*i)])
		b.SampleCount = append(b.SampleCount, sCount)
		b.SampleOffset = append(b.SampleOffset, int32(sOffset))
	}

	return box.FullRead(len(src))
}
