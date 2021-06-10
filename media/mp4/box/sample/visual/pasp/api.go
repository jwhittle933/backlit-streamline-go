package pasp

import (
	"encoding/binary"
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	sample2 "github.com/jwhittle933/streamline/media/mp4/box/sample"
)

const (
	PASP string = "pasp"
)

type Box struct {
	sample2.PixelAspectRatio
	raw []byte
}

func New(i *box2.Info) box2.Boxed {
	return &Box{
		sample2.PixelAspectRatio{
			Box:      base2.Box{BoxInfo: i},
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
	return b.Info().String() + fmt.Sprintf(
		", horiz_spacing=%d, vert_spacing=%d",
		b.HSpacing,
		b.VSpacing,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	b.HSpacing = binary.BigEndian.Uint32(src[0:4])
	b.VSpacing = binary.BigEndian.Uint32(src[4:8])
	b.raw = src

	return len(src), nil
}
