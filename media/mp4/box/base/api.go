package base

import (
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
)

type Box struct {
	BoxInfo *box2.Info
}

// Info satisfies the box.Informed interface
func (b *Box) Info() *box2.Info {
	return b.BoxInfo
}
