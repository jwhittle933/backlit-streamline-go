// Package tfhd (Track Fragment Header)
package tfhd

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/fullbox"
)

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
	fullbox.Box
	TrackID                uint32
	BaseDataOffset         uint64
	SampleDescriptionIndex uint32
	DefaultSampleDuration  uint32
	DefaultSampleSize      uint32
	DefaultSampleFlags     uint32
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
	}
}

func (Box) Type() string {
	return TFHD
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s, version=%d, flags=%d, trackid=%d, basedataoffset=%d, sampledescriptionindex=%d, defaultduration=%d, defaultsize=%d, defaultflags=%d",
		b.Info(),
		b.Version,
		b.Flags,
		b.TrackID,
		b.BaseDataOffset,
		b.SampleDescriptionIndex,
		b.DefaultSampleDuration,
		b.DefaultSampleSize,
		b.DefaultSampleFlags,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)
	b.TrackID = sr.Uint32()

	if b.Flags&BaseDataOffsetPreset != 0 {
		b.BaseDataOffset = sr.Uint64()
	}

	if b.Flags&SampleDescriptionIndexPresent != 0 {
		b.SampleDescriptionIndex = sr.Uint32()
	}

	if b.Flags&DefaultSampleDurationPresent != 0 {
		b.DefaultSampleDuration = sr.Uint32()
	}

	if b.Flags&DefaultSampleSizePresent != 0 {
		b.DefaultSampleSize = sr.Uint32()
	}

	if b.Flags&DefaultSampleFlagsPresent != 0 {
		b.DefaultSampleFlags = sr.Uint32()
	}

	return box.FullRead(len(src))
}
