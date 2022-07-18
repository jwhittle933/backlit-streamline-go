// Package idat (Image Data) for meta-data image files
package idat

import (
	"fmt"
	"github.com/jwhittle933/streamline/media/mpeg/base"
	box2 "github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	IDAT string = "idat"
)

type Box struct {
	base.Box
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base.Box{BoxInfo: i}}
}

func (Box) Type() string {
	return IDAT
}

func (b Box) String() string {
	return fmt.Sprintf(
		"[%s] hex=%s, offset=%d, size=%d, header=%d",
		b.BoxInfo.Type.String(),
		b.BoxInfo.Type.HexString(),
		b.BoxInfo.Offset,
		b.BoxInfo.Size,
		b.BoxInfo.HeaderSize,
	)
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	// iteratively parse children
	return len(src), nil
}
