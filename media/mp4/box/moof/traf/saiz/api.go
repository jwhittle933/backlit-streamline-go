// Package saiz (Sample Auxiliary Information Sizes)
package saiz

const (
	SAIZ string = "saiz"
)

type Box struct {
	AuxInfoType           [4]byte
	AuxInfoTypeParameter  uint32
	DefaultSampleInfoSize uint8
	SampleCount           uint32
	SampleInfoSize        []uint8
}

func (Box) Type() string {
	return SAIZ
}
