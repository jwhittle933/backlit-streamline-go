// Package smhd (Subtitle Media Header)
package sthd

import (
	"github.com/jwhittle933/streamline/media/mpeg/base"
	box2 "github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	STHD string = "sthd"
)

type Box struct {
	base.Box
	Balance   int16
	_reserved uint16
}

func New(i *box2.Info) box2.Boxed {
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
