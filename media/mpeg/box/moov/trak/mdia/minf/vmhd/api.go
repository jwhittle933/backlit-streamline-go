// Package vmhd (Video Media Header)
package vmhd

import (
	"github.com/jwhittle933/streamline/media/mpeg/base"
	box2 "github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	VMHD string = "vmhd"
)

type Box struct {
	base.Box
	GraphicsMode uint16
	OpColor      [3]uint16
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base.Box{BoxInfo: i}, 0, [3]uint16{}}
}

func (Box) Type() string {
	return VMHD
}

func (b *Box) String() string {
	return b.Info().String()
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
