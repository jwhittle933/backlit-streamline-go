// Package trik (Trick Play)
package trik

import (
	"fmt"
	"github.com/jwhittle933/streamline/media/mpeg/base"
	box2 "github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	TRIK string = "trik"
)

// Box is ISOBMFF mdat box type
type Box struct {
	base.Box
	Data []byte
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base.Box{BoxInfo: i}, []byte{}}
}

func (Box) Type() string {
	return TRIK
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
