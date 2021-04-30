package free

import "github.com/jwhittle933/streamline/pkg/media/mp4/box"

const (
	FREE string = "free"
)

// Box satisfies the box.Boxed interface
type Box struct {
	BoxInfo *box.Info
	Data []uint8
}

// New satisfies the mp4.BoxFactory function
func New(i *box.Info) box.Boxed {
	return &Box{BoxInfo: i}
}

// Type satisfies the box.Typed interface
func (Box) Type() string {
	return FREE
}

// Info satisfies the box.Informed interface
func (b *Box) Info() *box.Info {
	return b.BoxInfo
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
