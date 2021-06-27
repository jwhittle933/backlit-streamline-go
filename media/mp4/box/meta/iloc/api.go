// Package iloc (Item Location), for XML references to mandatory images, etc
package iloc

import (
	"fmt"
	"github.com/jwhittle933/streamline/media/mp4/base"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
)

const (
	ILOC string = "iloc"
)

type Box struct {
	base.Box
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base.Box{BoxInfo: i}}
}

func (Box) Type() string {
	return ILOC
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
