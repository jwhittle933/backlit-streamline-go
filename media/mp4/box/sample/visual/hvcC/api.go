package hvcC

import (
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	hvc2 "github.com/jwhittle933/streamline/media/mp4/box/sample/visual/hvc"
)

const (
	HVCC = "hvcC"
)

type Box struct {
	base2.Box
	hvc2.DecoderConfig
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, hvc2.New()}
}

func (b Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (Box) Type() string {
	return HVCC
}
