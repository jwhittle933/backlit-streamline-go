// Package frma (Original Format)
package frma

import (
	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	FRMA string = "frma"
)

// Box is ISOBMFF frma box type
type Box struct {
	base.Box
	DataFormat [4]byte
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, [4]byte{}}
}

func (Box) Type() string {
	return FRMA
}

func (b *Box) String() string {
	return b.Info().String()
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
