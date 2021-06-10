// Package hdlr (Handler) declares the media handler type
package hdlr

import (
	"encoding/binary"
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	HDLR string = "hdlr"
)

// Box is trak/mdia/hdlr box type
type Box struct {
	base2.Box
	Version uint8
	Flags   uint32
	// PreDefined: component_type of QuickTime
	// pre_defined of ISO-14496 always has 0
	// component_type has mhlr or dhlr
	Predefined  uint32
	HandlerType [4]byte
	Reserved    [3]uint64
	Name        string
	raw         []byte
}

func New(i *box2.Info) box2.Boxed {
	return &Box{
		base2.Box{BoxInfo: i},
		0,
		0,
		0,
		[4]byte{},
		[3]uint64{},
		"",
		make([]byte, 0),
	}
}

func (Box) Type() string {
	return HDLR
}

func (b *Box) String() string {
	return b.Info().String() + fmt.Sprintf(
		", version=%d, flags=%d, predefined=%d, handlertype=%s, name=%s",
		b.Version,
		b.Flags,
		b.Predefined,
		b.HandlerType,
		b.Name,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	b.Version = src[0]
	b.Flags = binary.BigEndian.Uint32([]byte{0x00, src[1], src[2], src[3]})

	b.raw = src[4:]
	b.Predefined = binary.BigEndian.Uint32(b.raw[0:4])
	copy(b.HandlerType[:], b.raw[4:8])
	// skip 12 reserved

	b.Name = string(b.raw[20:])
	return len(src), nil
}
