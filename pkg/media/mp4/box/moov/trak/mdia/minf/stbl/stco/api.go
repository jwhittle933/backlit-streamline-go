// Package stco (Chunk Offset)
package stco

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	STCO string = "stco"
)

// Box (stco) stores meta-data about video frames
type Box struct {
	base.Box
	EntryCount  uint32
	ChunkOffset []uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, 0, make([]uint32, 0)}
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
