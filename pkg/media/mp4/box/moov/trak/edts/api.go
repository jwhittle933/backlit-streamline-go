// Package edts (Edit Box)
package edts

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	EDTS string = "edts"
)

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
