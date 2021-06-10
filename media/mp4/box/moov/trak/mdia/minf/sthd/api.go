// Package smhd (Subtitle Media Header)
package sthd

import (
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	STHD string = "sthd"
)

type Box struct {
	base2.Box
	Balance   int16
	_reserved uint16
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, 0, 0}
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
