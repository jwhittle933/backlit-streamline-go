package urn

import (
	"encoding/binary"
	"fmt"
	"github.com/jwhittle933/streamline/media/mpeg/base"
	box2 "github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	URN           string = "urn "
	SelfContained uint   = 0x000001
)

type Box struct {
	base.Box
	Version  uint8
	Flags    uint32
	Name     string
	Location string
	raw      []byte
}

func (Box) Type() string {
	return URN
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base.Box{BoxInfo: i}, 0, 0, "", "", make([]byte, 0)}
}

func (b *Box) String() string {
	return b.Info().String() + fmt.Sprintf(
		", version=%d, flags=%d, name=%s, location=%s",
		b.Version,
		b.Flags,
		b.Name,
		b.Location,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	b.Version = src[0]
	b.Flags = binary.BigEndian.Uint32([]byte{0x00, src[1], src[2], src[3]})
	b.raw = src[4:]

	firstTerminator := len(b.raw)
	for i := 0; i < len(b.raw); i++ {
		if b.raw[i] == 0 {
			firstTerminator = i
		}
	}

	b.Name = string(b.raw[0:firstTerminator])
	b.Location = string(b.raw[firstTerminator:])

	return len(src), nil
}
