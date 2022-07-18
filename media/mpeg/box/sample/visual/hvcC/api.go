package hvcC

import (
	"fmt"
	"github.com/jwhittle933/streamline/media/mpeg/base"
	box2 "github.com/jwhittle933/streamline/media/mpeg/box"
	hvc2 "github.com/jwhittle933/streamline/media/mpeg/box/sample/visual/hvc"
)

const (
	HVCC = "hvcC"
)

type Box struct {
	base.Box
	hvc2.DecoderConfig
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base.Box{BoxInfo: i}, hvc2.New()}
}

func (b Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (Box) Type() string {
	return HVCC
}
