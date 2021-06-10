// Package frma (Original Format)
package frma

import (
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	FRMA string = "frma"
)

// Box is ISOBMFF frma box type
type Box struct {
	base2.Box
	DataFormat [4]byte
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, [4]byte{}}
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
