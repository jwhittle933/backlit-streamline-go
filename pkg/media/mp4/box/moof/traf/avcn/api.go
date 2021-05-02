// Package avcn (AVC NAL Unit Storage)
package avcn

import (
	"fmt"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	AVCN string = "avcn"
)

// Box is ISOBMFF mdat box type
type Box struct {
	base.Box
	Data    []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, []byte{}}
}

func (Box) Type() string {
	return AVCN
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
