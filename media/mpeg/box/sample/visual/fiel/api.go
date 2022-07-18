package fiel

import (
	"fmt"

	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	FIEL = "fiel"
)

type Box struct {
	base.Box
	raw []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, make([]byte, 0)}
}

func (Box) Type() string {
	return FIEL
}

func (b Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (b *Box) Write(src []byte) (int, error) {
	b.raw = src
	return box.FullRead(len(src))
}
