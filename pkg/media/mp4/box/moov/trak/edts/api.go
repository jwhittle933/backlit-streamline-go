// Package edts (Edit Box)
// See: https://github.com/itsjamie/bmff/blob/b0eabe94928cc481a1b373fc5e920257ab464912/edts.go
package edts

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/children"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/edts/elst"
)

const (
	EDTS string = "edts"
)

var Children = children.Registry{
	elst.ELST: elst.New,
}

// Box is ISOBMFF edts box type
type Box struct {
	base.Box
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}}
}

func (Box) Type() string {
	return EDTS
}

func (b Box) String() string {
	return b.Info().String()
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
