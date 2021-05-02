// Package tfhd (Track Fragment Header)
package tfhd

const (
	TFHD                          string = "tfhd"
	BaseDataOffsetPreset                 = 0x000001
	SampleDescriptionIndexPresent        = 0x000002
	DefaultSampleDurationPresent         = 0x000008
	DefaultSampleSizePresent             = 0x000010
	DefaultSampleFlagsPresent            = 0x000020
	DurationIsEmpty                      = 0x010000
	DefaultBaseIsMOOF                    = 0x020000
)

type Box struct {
	TrackID                uint32
	BaseDataOffset         uint64
	SampleDescriptionIndex uint32
	DefaultSampleDuration  uint32
	DefaultSampleSize      uint32
	DefaultSampleFlags     uint32
}

func (Box) Type() string {
	return TFHD
}
