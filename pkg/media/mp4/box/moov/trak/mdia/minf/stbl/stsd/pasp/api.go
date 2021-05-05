package pasp

import (
	"encoding/binary"
	"fmt"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/sample"
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
	fmt.Println("PASP Stringer")
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
