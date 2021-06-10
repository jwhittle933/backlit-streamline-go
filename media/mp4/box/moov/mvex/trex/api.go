// Package trex (Track Extends)
package trex

const (
	TREX string = "trex"
)

type Box struct {
	TrackID                  uint32
	DefaultSampleDescription uint32
	DefaultSampleDuration    uint32
	DefaultSampleSize        uint32
	DefaultSampleFlags       uint32
}

func (Box) Type() string {
	return TREX
}
