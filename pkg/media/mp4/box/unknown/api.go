package unknown

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
)

const (
	UNKNOWN = "unkn"
)

type Box struct {
	BoxInfo *box.Info
	Data    []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{BoxInfo: i}
}

func (Box) Type() string {
	return UNKNOWN
}

func (b *Box) Info() *box.Info {
	return b.BoxInfo
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}

