package tkhd

const (
	TKHD string = "tkhd"
)

type Measurement uint32

type Box struct {
	CreationTimeV0     uint32
	ModificationTimeV0 uint32
	CreationTimeV1     uint64
	ModificationTimeV1 uint64
	TrackID        uint32
	_reserved0     uint32
	DurationV0     uint32
	DurationV1     uint64
	_reserved1     [2]uint32
	Layer          int16
	AlternateGroup int16
	Volume         int16
	_reserved2     uint16
	Matrix         [9]int32
	Width          Measurement
	Height         Measurement
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
