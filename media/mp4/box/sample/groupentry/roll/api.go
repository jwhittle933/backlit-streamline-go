package roll

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	ROLL = "roll"
)

type Box struct {
	base.Box
	RollDistance uint16
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
	}
}

func (Box) Type() string {
	return ROLL
}

func (b Box) String() string {
	return fmt.Sprintf("%s, rolldistance=%d", b.Info(), b.RollDistance)
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.RollDistance = sr.Uint16()

	return box.FullRead(len(src))
}
