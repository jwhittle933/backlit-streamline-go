// Package mdat for ISO BMFF Media Data box
// "The media data box (‘mdat’) is merely one possible location [for the data],
// and looked at by itself, it can only be considered an un‐ordered bag of
// un‐identifiable bits."
package mdat

import (
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	MDAT string = "mdat"
)

// Box is ISOBMFF mdat box type
type Box struct {
	base2.Box
	Data []byte
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, []byte{}}
}

func (Box) Type() string {
	return MDAT
}

func (b Box) String() string {
	return fmt.Sprintf(
		"[%s] hex=%s, offset=%d, size=%d, header=%d",
		string(b.BoxInfo.Type.String()),
		b.BoxInfo.Type.HexString(),
		b.BoxInfo.Offset,
		b.BoxInfo.Size,
		b.BoxInfo.HeaderSize,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	b.Data = src
	return len(src), nil
}
