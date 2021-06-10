// Package stco (Chunk Offset)
package stco

import (
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	STCO string = "stco"
)

// Box (stco) stores meta-data about video frames
type Box struct {
	base2.Box
	EntryCount  uint32
	ChunkOffset []uint32
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, 0, make([]uint32, 0)}
}

func (Box) Type() string {
	return STCO
}

func (b *Box) String() string {
	return b.Info().String()
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
