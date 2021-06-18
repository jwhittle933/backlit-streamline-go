// Package mdhd (Media Header)
package mdhd

import (
	"encoding/binary"
	"fmt"
	"github.com/jwhittle933/streamline/media/mp4/fullbox"

	"github.com/jwhittle933/streamline/media/mp4/box"
)

const (
	MDHD string = "mdhd"
)

// Box is ISOBMFF mdhd box type
type Box struct {
	fullbox.Box
	CreationTime     uint64
	ModificationTime uint64
	Timescale        uint32
	Duration         uint64
	Pad              bool
	Language         [2]byte
	LanguageCode     string
	Predefined       uint16
	raw              []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{
		*fullbox.New(i),
		0,
		0,
		0,
		0,
		false,
		[2]byte{},
		"",
		0,
		make([]byte, 0),
	}
}

func (Box) Type() string {
	return MDHD
}

func (b *Box) String() string {
	return b.Info().String() + fmt.Sprintf(
		", version=%d, flags=%d, creation=%d, modification=%d, timescale=%d, duration=%d, pad=%+v, lang=%d, language=%s",
		b.Version,
		b.Flags,
		b.CreationTime,
		b.ModificationTime,
		b.Timescale,
		b.Duration,
		b.Pad,
		b.Language,
		b.LanguageCode,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	b.Version = src[0]
	b.Flags = binary.BigEndian.Uint32([]byte{0x00, src[1], src[2], src[3]})

	b.raw = src[4:]

	offset := 0
	if b.Version == 0 {
		b.CreationTime = uint64(binary.BigEndian.Uint32(b.raw[0:4]))
		b.ModificationTime = uint64(binary.BigEndian.Uint32(b.raw[4:8]))
		b.Timescale = binary.BigEndian.Uint32(b.raw[8:12])
		b.Duration = uint64(binary.BigEndian.Uint32(b.raw[12:16]))

		offset = 16
	} else if b.Version == 1 {
		b.CreationTime = binary.BigEndian.Uint64(b.raw[0:8])
		b.ModificationTime = binary.BigEndian.Uint64(b.raw[8:16])
		b.Timescale = binary.BigEndian.Uint32(b.raw[16:20])
		b.Duration = uint64(binary.BigEndian.Uint32(b.raw[20:28]))

		offset = 28
	}

	copy(b.Language[:], b.raw[offset:offset+2])
	lang := binary.BigEndian.Uint16(b.Language[:])
	b.LanguageCode = string([]byte{
		uint8(lang&0x7c00>>10) + 0x60,
		uint8(lang&0x03E0>>5) + 0x60,
		uint8(lang&0x001F) + 0x60,
	})

	b.Predefined = binary.BigEndian.Uint16(b.raw[offset+2 : offset+4])
	return len(src), nil
}
