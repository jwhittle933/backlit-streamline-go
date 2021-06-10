package udta

import (
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	UDTA string = "udta"
)

// Box is ISOBMFF udta box type
type Box struct {
	base2.Box
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}}
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
