package fullbox

import (
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	FlagsMask = 0x00ffffff
)

type Box struct {
	base.Box
	Version byte
	Flags   uint32
}

func New(i *box.Info) *Box {
	return &Box{base.Box{BoxInfo: i}, 0, 0}
}

func (b *Box) WriteVersionAndFlags(sr *slicereader.Reader) {
	versionAndFlags := sr.Uint32()
	b.Version = byte(versionAndFlags >> 24)
	b.Flags = versionAndFlags & FlagsMask
}
