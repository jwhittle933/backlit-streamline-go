package nmhd

import (
	"fmt"
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/fullbox"
)

const (
	NMHD = "nmhd"
)

type Box struct {
	fullbox.Box
}

func New(i *box.Info) box.Boxed {
	return &Box{*fullbox.New(i)}
}

func (Box) Type() string {
	return NMHD
}

func (b Box) String() string {
	return fmt.Sprintf("%s, version=%d, flags=%d", b.Info(), b.Version, b.Flags)
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)
	return box.FullRead(len(src))
}
