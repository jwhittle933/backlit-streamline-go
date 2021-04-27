package avc1

const (
	AVC1 string = "avc1"
)

type Box struct {
	AuxInfoType           [4]byte
	AuxInfoTypeParameter  uint32
	DefaultSampleInfoSize uint8
	SampleCount           uint32
	SampleInfoSize        []uint8
}

func (Box) Type() string {
	return AVC1
}
