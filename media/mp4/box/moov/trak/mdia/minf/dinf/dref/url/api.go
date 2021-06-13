package url

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	URL           string = "url "
	SelfContained uint32 = 0x000001
)

type Box struct {
	base.Box
	Version  uint8
	Flags    uint32
	Location string
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, 0, 0, ""}
}

func (Box) Type() string {
	return URL
}

func (b *Box) String() string {
	return b.Info().String() + fmt.Sprintf(
		", version=%d, flags=%d, location=%s",
		b.Version,
		b.Flags,
		b.Location,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	versionAndFlags := sr.Uint32()

	b.Version = byte(versionAndFlags >> 24)
	b.Flags = versionAndFlags & box.FlagsMask

	if b.Flags&SelfContained > 0 {
		b.Location = sr.String(sr.Length()-1)
	}

	return len(src), nil
}
