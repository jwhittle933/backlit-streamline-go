// Package tkhd (Track Header)
package tkhd

import (
	"encoding/binary"
	"fmt"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	TKHD string = "tkhd"
)

type Measurement uint32

type Box struct {
	base.Box
	Version          uint8
	Flags            uint32
	CreationTime     uint64
	ModificationTime uint64
	TrackID          uint32
	_reserved0       uint32
	Duration         uint64
	_reserved1       [2]uint32
	Layer            int16
	AlternateGroup   int16
	Volume           int16
	_reserved2       uint16
	Matrix           [9]int32
	Width            Measurement
	Height           Measurement
	raw              []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		[2]uint32{},
		0,
		0,
		0,
		0,
		[9]int32{},
		0,
		0,
		make([]byte, 0),
	}
}

func (Box) Type() string {
	return TKHD
}

func (m Measurement) Float64() float64 {
	return float64(m) / (1 << 16)
}

func (m Measurement) UInt16() uint16 {
	return uint16(m >> 16)
}

func (b Box) String() string {
	return b.Info().String() + fmt.Sprintf(
		", version=%d, flags=%d, creation=%d, modification=%d",
		b.Version,
		b.Flags,
		b.CreationTime,
		b.ModificationTime,
	)
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	b.raw = src
	b.Version = src[0]
	b.Flags = binary.BigEndian.Uint32([]byte{0x00, b.raw[1], b.raw[2], b.raw[3]})
	b.raw = src[4:]

	var offset int
	if b.Version == 0 {
		b.CreationTime = uint64(binary.BigEndian.Uint32(b.raw[0:4]))
		b.ModificationTime = uint64(binary.BigEndian.Uint32(b.raw[4:8]))
		b.TrackID = binary.BigEndian.Uint32(b.raw[8:12])
		// 12:16 reserved
		b.Duration = uint64(binary.BigEndian.Uint32(b.raw[16:20]))
		offset = 20
	} else if b.Version == 1 {
		b.CreationTime = binary.BigEndian.Uint64(b.raw[0:8])
		b.ModificationTime = binary.BigEndian.Uint64(b.raw[8:16])
		b.TrackID = binary.BigEndian.Uint32(b.raw[16:20])
		// 20:24 reserved
		b.Duration = uint64(binary.BigEndian.Uint32(b.raw[24:28]))
		offset = 20
	}
	offset += 8

	b.Layer = int16(binary.BigEndian.Uint16(b.raw[offset : offset+2]))
	offset += 2

	b.AlternateGroup = int16(binary.BigEndian.Uint16(b.raw[offset : offset+2]))
	offset += 2

	b.Volume = int16(binary.BigEndian.Uint16(b.raw[offset : offset+2]))
	offset += 4 // plus 2 reserved bytes

	for i := 0; i < 9; i++ {
		b.Matrix[i] = int32(binary.BigEndian.Uint32(b.raw[offset+i : offset+i+4]))
	}
	offset += 36

	b.Width = Measurement(binary.BigEndian.Uint32(b.raw[offset : offset+4]))
	offset += 4

	b.Height = Measurement(binary.BigEndian.Uint32(b.raw[offset : offset+4]))
	offset += 4

	return len(src), nil
}
