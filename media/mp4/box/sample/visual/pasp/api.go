package pasp

import (
	"fmt"
	"github.com/jwhittle933/streamline/bits/slicereader"

	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	PASP string = "pasp"
)

type Box struct {
	base.Box
	HSpacing uint32
	VSpacing uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{
		Box:      base.Box{BoxInfo: i},
		HSpacing: 0,
		VSpacing: 0,
	}
}

func (Box) Type() string {
	return PASP
}

func (b *Box) String() string {
	return fmt.Sprintf("%s, horiz_spacing=%d, vert_spacing=%d", b.Info(), b.HSpacing, b.VSpacing)
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.HSpacing = sr.Uint32()
	b.VSpacing = sr.Uint32()

	return len(src), nil
}
