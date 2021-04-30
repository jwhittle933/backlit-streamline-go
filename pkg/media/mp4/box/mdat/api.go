package mdat

import "github.com/jwhittle933/streamline/pkg/media/mp4/box"

const (
	MDAT string = "mdat"
)

// Box is ISOBMFF mdat box type
type Box struct {
	BoxInfo *box.Info
	Data    []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{BoxInfo: i}
}

func (Box) Type() string {
	return MDAT
}

func (b *Box) Info() *box.Info {
	return b.BoxInfo
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
