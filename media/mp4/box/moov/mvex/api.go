// Package mvex (Movie Extends)
package mvex

import (
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	children2 "github.com/jwhittle933/streamline/media/mp4/box/children"
)

const (
	MVEX string = "mvex"
)

var (
	Children = children2.Registry{}
)

type Box struct {
	base2.Box
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}}
}

func (Box) Type() string {
	return MVEX
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
