package free

import (
	"fmt"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	FREE string = "free"
)

// Box satisfies the box.Boxed interface
type Box struct {
	base.Box
	Data []uint8
}

// New satisfies the mp4.BoxFactory function
func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, []uint8{}}
}

// Type satisfies the box.Typed interface
func (Box) Type() string {
	return FREE
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

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
