// Package mvhd (Movie Header)
package mvhd

import (
	"encoding/binary"
	"fmt"
	"github.com/jwhittle933/streamline/media/mpeg/base"
	box2 "github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	MVHD string = "mvhd"
)

type Box struct {
	base.Box
	Version          uint8
	Flags            uint32
	CreationTime     uint64
	ModificationTime uint64
	Timescale        uint32
	Duration         uint64
	Rate             Rate
	Volume           uint16
	_reserved        uint16
	_reserved2       [2]uint32
	Matrix           [9]int32
	Predefined       []byte
	NextTrackID      uint32
	raw              []byte
}

type Rate uint32

func New(i *box2.Info) box2.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		[2]uint32{},
		[9]int32{},
		make([]byte, 0),
		0,
		make([]byte, 0),
	}
}

func (Box) Type() string {
	return MVHD
}

func (b *Box) Write(src []byte) (int, error) {
	b.raw = src
	b.Version = src[0]
	b.Flags = binary.BigEndian.Uint32([]byte{0x00, b.raw[1], b.raw[2], b.raw[3]})

	b.raw = src[4:]

	var offset int
	if b.Version == 0 {
		b.CreationTime = uint64(binary.BigEndian.Uint32(b.raw[0:4]))
		b.ModificationTime = uint64(binary.BigEndian.Uint32(b.raw[4:8]))
		b.Timescale = binary.BigEndian.Uint32(b.raw[8:12])
		b.Duration = uint64(binary.BigEndian.Uint32(b.raw[12:16]))
		offset += 16
	} else if b.Version == 1 {
		b.CreationTime = binary.BigEndian.Uint64(b.raw[0:8])
		b.ModificationTime = binary.BigEndian.Uint64(b.raw[8:16])
		b.Timescale = binary.BigEndian.Uint32(b.raw[16:20])
		b.Duration = binary.BigEndian.Uint64(b.raw[20:28])
		offset += 28
	}

	b.Rate = Rate(binary.BigEndian.Uint32(b.raw[offset : offset+4]))
	offset += 4

	b.Volume = binary.BigEndian.Uint16(b.raw[offset : offset+2])
	offset += 2

	b._reserved = binary.BigEndian.Uint16(b.raw[offset : offset+10])
	offset += 10

	for i := 0; i < 9; i++ {
		b.Matrix[i] = int32(binary.BigEndian.Uint32(b.raw[offset+i : offset+i+4]))
	}
	offset += 9 * 4

	// rethink []byte. Should be [6]int32
	copy(b.Predefined, src[offset:offset+24])
	offset += 24

	b.NextTrackID = binary.BigEndian.Uint32(src[offset : offset+4])

	return len(src), nil
}

func (b *Box) String() string {
	return b.Info().String() + fmt.Sprintf(
		", version=%d, flags=%d, creation=%d, modification=%d, timescale=%d, duration=%d, rate=%d, volume=%d, matrix=%+v, predefined=%+v, next_track_id=%d",
		b.Version,
		b.Flags,
		b.CreationTime,
		b.ModificationTime,
		b.Timescale,
		b.Duration,
		b.Rate,
		b.Volume,
		b.Matrix,
		b.Predefined,
		b.NextTrackID,
	)
}

func (r Rate) Float64() float64 {
	return float64(r) / (1 << 16)
}

func (r Rate) Int16() int16 {
	return int16(r >> 16)
}
