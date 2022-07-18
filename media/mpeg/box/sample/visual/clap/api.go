package clap

import (
	"fmt"
	slicereader2 "github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/base"
	box2 "github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	CLAP = "clap"
)

type Box struct {
	base.Box
	CleanApertureWidthN  uint32
	CleanApertureWidthD  uint32
	CleanApertureHeightN uint32
	CleanApertureHeightD uint32
	HorizOffN            uint32
	HorizOffD            uint32
	VertOffN             uint32
	VertOffD             uint32
}

func New(i *box2.Info) box2.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
	}
}

func (b Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (Box) Type() string {
	return CLAP
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader2.New(src)

	b.CleanApertureWidthN = sr.Uint32()
	b.CleanApertureHeightD = sr.Uint32()
	b.CleanApertureHeightN = sr.Uint32()
	b.CleanApertureHeightD = sr.Uint32()
	b.HorizOffN = sr.Uint32()
	b.HorizOffD = sr.Uint32()
	b.VertOffN = sr.Uint32()
	b.VertOffD = sr.Uint32()

	return box2.FullRead(len(src))
}
