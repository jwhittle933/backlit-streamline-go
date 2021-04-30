package free

import "github.com/jwhittle933/streamline/pkg/media/mp4/box"

const (
	FREE string = "free"
)

type Box struct {
	BoxInfo *box.Info
	Data []uint8
}

func New(i *box.Info) box.Boxed {
	return &Box{BoxInfo: i}
}

func (Box) Type() string {
	return FREE
}

func (b *Box) Info() *box.Info {
	return b.BoxInfo
}
