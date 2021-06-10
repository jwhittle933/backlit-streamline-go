// Package meta for multi-track required meta-data
package meta

import (
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	children2 "github.com/jwhittle933/streamline/media/mp4/box/children"
)

const (
	META string = "meta"
)

var (
	Children = children2.Registry{}
)

// Box is ISO BMFF meta box type
type Box struct {
	base2.Box
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}}
}

func (Box) Type() string {
	return META
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}

func (b *Box) String() string {
	return fmt.Sprintf(
		"[%s] hex=%s, offset=%d, size=%d, header=%d",
		string(b.BoxInfo.Type.String()),
		b.BoxInfo.Type.HexString(),
		b.BoxInfo.Offset,
		b.BoxInfo.Size,
		b.BoxInfo.HeaderSize,
	)
}
