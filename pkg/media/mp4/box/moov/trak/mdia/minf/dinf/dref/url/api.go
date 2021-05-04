package url

import (
	"encoding/binary"
	"fmt"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
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
	raw      []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, 0, 0, "", make([]byte, 0)}
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
	b.Version = src[0]
	b.Flags = binary.BigEndian.Uint32([]byte{0x00, src[1], src[2], src[3]})

	if b.Flags&SelfContained > 0 {
		b.Location = string(src[4:])
	}

	b.raw = src
	return len(src), nil
}
