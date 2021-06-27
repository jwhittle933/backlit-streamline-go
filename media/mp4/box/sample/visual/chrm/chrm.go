package chrm

import (
	"fmt"

	"github.com/jwhittle933/streamline/media/mp4/base"
	"github.com/jwhittle933/streamline/media/mp4/box"
)

const (
	CHRM = "chrm"
)

type Box struct {
	base.Box
	raw []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, nil}
}

func (Box) Type() string {
	return CHRM
}

func (b Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (b *Box) Write(src []byte) (int, error) {
	b.raw = src
	return box.FullRead(len(src))
}
