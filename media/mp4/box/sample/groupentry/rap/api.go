package rap

import (
	"fmt"

	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	RAP = "rap "
)

type Box struct {
	base.Box
	NumLeadingSamplesKnown uint8
	NumLeadingSamples      uint8
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
	}
}

func (Box) Type() string {
	return RAP
}

func (b Box) String() string {
	return fmt.Sprintf("%s, leading_samples_known=%d, leading_samples=%d", b.Info(), b.NumLeadingSamplesKnown, b.NumLeadingSamples)
}

func (b *Box) Write(src []byte) (int, error) {
	b.NumLeadingSamplesKnown = src[0] >> 7
	b.NumLeadingSamples = src[0] & 0x7F

	return box.FullRead(len(src))
}
