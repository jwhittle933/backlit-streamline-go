package moov

import "github.com/jwhittle933/streamline/pkg/media/mp4/box"

const (
	MOOV string = "moov"
)

type Box struct {
	BoxInfo *box.Info
}

func New(i *box.Info) box.Boxed {
	return &Box{BoxInfo: i}
}

func (Box) Type() string {
	return MOOV
}

func (b *Box) Info() *box.Info {
	return b.BoxInfo
}
