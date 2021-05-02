// Package stsz (Sample Size)
package stsz

const (
	STSZ string = "stsz"
)

type Box struct {
	SampleSize  uint32
	SampleCount uint32
	EntrySize   []uint32
}

func (Box) Type() string {
	return STSZ
}
