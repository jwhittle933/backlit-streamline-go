package co64

import (
	"fmt"
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	CO64 string = "co64"
)

type Box struct {
	base.Box
	Version     byte
	Flags       uint32
	ChunkOffset []uint64
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		make([]uint64, 0),
	}
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s, version=%d, flags=%d, chuckoffset=%+v",
		b.Info(),
		b.Version,
		b.Flags,
		b.ChunkOffset,
	)
}

func (Box) Type() string {
	return CO64
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	versionAndFlags := sr.Uint32()
	nrEntries := sr.Uint32()

	b.Version = byte(versionAndFlags >> 24)
	b.Flags = versionAndFlags & box.FlagsMask

	b.ChunkOffset = make([]uint64, nrEntries)
	for i := uint32(0); i < nrEntries; i++ {
		b.ChunkOffset[i] = sr.Uint64()
	}

	return box.FullRead(len(src))
}
