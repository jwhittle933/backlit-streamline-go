// Package trex (Track Extends)
package trex

import (
	"fmt"
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/fullbox"
)

const (
	TREX string = "trex"
)

type Box struct {
	fullbox.Box
	Version                       byte
	Flags                         uint32
	TrackID                       uint32
	DefaultSampleDescriptionIndex uint32
	DefaultSampleDuration         uint32
	DefaultSampleSize             uint32
	DefaultSampleFlags            uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{
		*fullbox.New(i),
		0,
		0,
		0,
		0,
		0,
		0,
		0,
	}
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s, version=%d, flags=%d, trackid=%d, index=%d, defaultduration=%d, defaultsize=%d, defaultflags=%d",
		b.Info(),
		b.Version,
		b.Flags,
		b.TrackID,
		b.DefaultSampleDescriptionIndex,
		b.DefaultSampleDuration,
		b.DefaultSampleSize,
		b.DefaultSampleFlags,
	)
}

func (b Box) Type() string {
	return TREX
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)
	b.TrackID = sr.Uint32()
	b.DefaultSampleDescriptionIndex = sr.Uint32()
	b.DefaultSampleDuration = sr.Uint32()
	b.DefaultSampleSize = sr.Uint32()
	b.DefaultSampleFlags = sr.Uint32()

	return box.FullRead(len(src))
}
