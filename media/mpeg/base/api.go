package base

import (
	"github.com/jwhittle933/streamline/media/mpeg/box"
)

type Box struct {
	BoxInfo *box.Info
}

// Info satisfies the box.Informed interface
func (b *Box) Info() *box.Info {
	return b.BoxInfo
}
