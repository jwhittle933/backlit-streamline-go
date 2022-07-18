package fullbox

import (
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
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
	version, flags := VersionAndFlags(sr.Uint32())
	b.Version = version
	b.Flags = flags
}

func VersionAndFlags(vf uint32) (byte, uint32) {
	return byte(vf >> 24), vf & FlagsMask
}
