// Package mdat for ISO BMFF Media Data box
// "The media data box (‘mdat’) is merely one possible location [for the data],
// and looked at by itself, it can only be considered an un‐ordered bag of
// un‐identifiable bits."
package mdat

import (
	"fmt"

	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	MDAT string = "mdat"
)

// Box is ISO BMFF mdat box type
type Box struct {
	base.Box
	Data []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, make([]byte, 0)}
}

func (Box) Type() string {
	return MDAT
}

func (b Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (b *Box) Write(src []byte) (int, error) {
	b.Data = src
	return box.FullRead(len(src))
}
