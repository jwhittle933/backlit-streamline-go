// Package mvhd (Movie Header)
package mvhd

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	MVHD string = "mvhd"
)

type Box struct {
	base.Box
	CreationTimeV0     uint32
	ModificationTimeV0 uint32
	CreationTimeV1     uint64
	ModificationTimeV1 uint64
	Timescale          uint32
	DurationV0         uint32
	DurationV1         uint64
	Rate               Rate
	Volume             int16
	_reserved          int16
	_reserved2         [2]uint32
	Matrix             [9]int32
	Predefined         [6]int32
	NextTrackID        uint32
}

type Rate int32

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		[2]uint32{}, [9]int32{}, [6]int32{}, 0,
	}
}

func (Box) Type() string {
	return MVHD
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}

func (b *Box) String() string {
	return ""
}

func (r Rate) Float64() float64 {
	return float64(r) / (1 << 16)
}

func (r Rate) Int16() int16 {
	return int16(r >> 16)
}
