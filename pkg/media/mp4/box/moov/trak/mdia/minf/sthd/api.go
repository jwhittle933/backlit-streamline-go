// Package smhd (Subtitle Media Header)
package sthd

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	STHD string = "sthd"
)

type Box struct {
	base.Box
	Balance   int16
	_reserved uint16
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, 0, 0}
}

func (Box) Type() string {
	return STHD
}

func (b *Box) String() string {
	return b.Info().String()
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
