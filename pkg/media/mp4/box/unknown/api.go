package unknown

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
)

const (
	UNKNOWN = "unkn"
)

type Box struct {
	BoxInfo *box.Info
}

func New(i *box.Info) box.Boxed {
	return Box{BoxInfo: i}
}

func (Box) Type() string {
	return UNKNOWN
}

func (b Box) Info() *box.Info {
	return b.BoxInfo
}
