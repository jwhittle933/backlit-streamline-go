package free

import (
	"fmt"

	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	FREE string = "free"
)

// Box satisfies the box.Boxed interface
type Box struct {
	base.Box
	Data []byte
}

// New satisfies the mpeg.BoxFactory function
func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, []uint8{}}
}

// Type satisfies the box.Typed interface
func (Box) Type() string {
	return FREE
}

func (b Box) String() string {
	return fmt.Sprintf("%s, data=%s", b.Info(), b.Data)
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	b.Data = src
	return len(src), nil
}
