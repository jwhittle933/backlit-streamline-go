// Package stco (Chunk Offset)
package stco

import (
	"encoding/binary"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/base"
	"github.com/jwhittle933/streamline/media/mp4/box"
)

const (
	STCO string = "stco"
)

// Box (stco) stores meta-data about video frames
type Box struct {
	base.Box
	Version     byte
	Flags       uint32
	EntryCount  uint32
	ChunkOffset []uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		make([]uint32, 0),
	}
}

func (Box) Type() string {
	return STCO
}

func (b *Box) String() string {
	return b.Info().String()
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	versionAndFlags := sr.Uint32()
	b.Version = byte(versionAndFlags >> 24)
	b.Flags = versionAndFlags & box.FlagsMask
	b.EntryCount = sr.Uint32()

	for i := uint32(0); i < b.EntryCount; i++ {
		b.ChunkOffset = append(b.ChunkOffset, binary.BigEndian.Uint32(src[(8+4*i):(12+4*i)]))
	}

	return box.FullRead(len(src))
}
