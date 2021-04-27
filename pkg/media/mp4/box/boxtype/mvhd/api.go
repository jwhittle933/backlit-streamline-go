package mvhd

const (
	MVHD string = "mvhd"
)

type Box struct {
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

func (Box) Type() string {
	return MVHD
}

func (r Rate) Float64() float64 {
	return float64(r) / (1 << 16)
}

func (r Rate) Int16() int16 {
	return int16(r >> 16)
}
