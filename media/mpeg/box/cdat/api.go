package cdat

import (
	"fmt"
	"github.com/jwhittle933/streamline/media/mpeg/base"
	box2 "github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	CDAT = "cdat"
)

type Box struct {
	base.Box
	Data []byte
}

func New(i *box2.Info) *Box {
	return &Box{base.Box{BoxInfo: i}, make([]byte, 0, 0)}
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s, datalength=%d",
		b.Info(),
		len(b.Data),
	)
}

func (Box) Type() string {
	return CDAT
}

func (b *Box) Write(src []byte) (int, error) {
	b.Data = src

	return len(src), nil
}
