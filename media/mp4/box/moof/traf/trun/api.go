// Package trun (Track Fragment Run)
package trun

const (
	TRUN string = "trun"
)

type Box struct {
	SampleCount      uint32
	DataOffset       int32
	FirstSampleFlags uint32
	Entries          []Entry
}

type Entry struct {
	SampleDuration                uint32
	SampleSize                    uint32
	SampleFlags                   uint32
	SampleCompositionTimeOffsetV0 uint32
	SampleCompositionTimeOffsetV1 int32
}

func (Box) Type() string {
	return TRUN
}
