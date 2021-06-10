// Package smhd (Sound Media Header)
package smhd

import (
	"encoding/binary"
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	SMHD string = "smhd"
)

type Box struct {
	base2.Box
	Version   uint8
	Flags     uint32
	Balance   int16
	_reserved uint16
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, 0, 0, 0, 0}
}

func (Box) Type() string {
	return SMHD
}

func (b *Box) String() string {
	return b.Info().String() + fmt.Sprintf(
		", version=%d, flags=%d, balance=%d",
		b.Version,
		b.Flags,
		b.Balance,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	b.Version = src[0]
	b.Flags = binary.BigEndian.Uint32([]byte{0x00, src[1], src[2], src[3]})

	b.Balance = int16(binary.BigEndian.Uint16(src[4:6]))
	b._reserved = binary.BigEndian.Uint16(src[6:8])

	return len(src), nil
}
