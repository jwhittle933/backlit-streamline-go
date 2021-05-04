package urn

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	URN string = "urn "
	SelfContained uint = 0x000001
)

type Box struct {
	base.Box
	Name     string
	Location string
}

func (Box) Type() string {
	return URN
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, "", ""}
}

func (b *Box) String() string {
	return b.Info().String()
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
