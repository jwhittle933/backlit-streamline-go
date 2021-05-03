package udta

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	UDTA string = "udta"
)

// Box is ISOBMFF udta box type
type Box struct {
	base.Box
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}}
}

type Udat3GPPString struct {
	Pad      bool
	Language [3]byte
	Data     []byte
}

func (Box) Type() string {
	return UDTA
}

func (b Box) String() string {
	return b.Info().String()
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
