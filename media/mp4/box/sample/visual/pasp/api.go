package pasp

import (
	"encoding/binary"
	"fmt"

	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
	"github.com/jwhittle933/streamline/media/mp4/box/sample"
)

const (
	PASP string = "pasp"
)

type Box struct {
	sample.PixelAspectRatio
	raw []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{
		sample.PixelAspectRatio{
			Box:      base.Box{BoxInfo: i},
			HSpacing: 0,
			VSpacing: 0,
		},
		make([]byte, 0),
	}
}

func (Box) Type() string {
	return PASP
}

func (b *Box) String() string {
	return fmt.Sprintf("%s, horiz_spacing=%d, vert_spacing=%d", b.Info(), b.HSpacing, b.VSpacing)
}

func (b *Box) Write(src []byte) (int, error) {
	b.HSpacing = binary.BigEndian.Uint32(src[0:4])
	b.VSpacing = binary.BigEndian.Uint32(src[4:8])
	b.raw = src

	return len(src), nil
}
