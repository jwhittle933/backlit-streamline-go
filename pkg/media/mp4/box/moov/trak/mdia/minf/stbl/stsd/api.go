// Package stsd (Sample Descriptions)
package stsd

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	STSD string = "stsd"
)

type Box struct {
	base.Box
	EntryCount uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, 0}
}

func (Box) Type() string {
	return STSD
}

func (b *Box) String() string {
	return b.Info().String()
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
