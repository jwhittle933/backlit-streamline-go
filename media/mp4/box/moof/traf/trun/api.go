// Package trun (Track Fragment Run)
package trun

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/fullbox"
)

const (
	TRUN string = "trun"
)

const (
	DataOffsetPresentFlag       uint32 = 0x01
	FirstSampleFlagsPresentFlag uint32 = 0x04
	SampleDurationPresentFlag   uint32 = 0x100
	SampleSizePresentFlag       uint32 = 0x200
	SampleFlagsPresentFlag      uint32 = 0x400
	SampleCTOPresentFlag        uint32 = 0x800
)

type Box struct {
	fullbox.Box
	SampleCount      uint32
	DataOffset       int32
	FirstSampleFlags uint32
	Samples          []Sample
	WriteOrderNr     uint32
}

type Sample struct {
	SampleDuration              uint32
	SampleSize                  uint32
	SampleFlags                 uint32
	SampleCompositionTimeOffset int32
}

func (s Sample) String() string {
	return fmt.Sprintf(
		"[\033[1;34msample\033[0m] duration=%d, size=%d, flags=%d, cto=%d",
		s.SampleDuration,
		s.SampleSize,
		s.SampleFlags,
		s.SampleCompositionTimeOffset,
	)
}

func New(i *box.Info) box.Boxed {
	return &Box{
		*fullbox.New(i),
		0,
		0,
		0,
		make([]Sample, 0),
		0,
	}
}

func (b Box) String() string {
	s := fmt.Sprintf(
		"%s, version=%d, flags=%d, samples=%d",
		b.Info(),
		b.Version,
		b.Flags,
		len(b.Samples),
	)

	for _, sa := range b.Samples {
		s += fmt.Sprintf("\n      %s", sa)
	}

	return s
}

func (Box) Type() string {
	return TRUN
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)
	b.SampleCount = sr.Uint32()
	b.Samples = make([]Sample, b.SampleCount)

	if b.Flags&DataOffsetPresentFlag != 0 {
		b.DataOffset = int32(sr.Uint32())
	}

	if b.Flags&FirstSampleFlagsPresentFlag != 0 {
		b.FirstSampleFlags = sr.Uint32()
	}

	for i := 0; i < len(b.Samples); i++ {
		s := Sample{}

		if b.Flags&SampleDurationPresentFlag != 0 {
			s.SampleDuration = sr.Uint32()
		}

		if b.Flags&SampleSizePresentFlag != 0 {
			s.SampleSize = sr.Uint32()
		}

		if b.Flags&SampleFlagsPresentFlag != 0 {
			s.SampleFlags = sr.Uint32()
		} else if b.Flags&FirstSampleFlagsPresentFlag != 0 {
			s.SampleFlags = b.FirstSampleFlags
		}

		if b.Flags&SampleCTOPresentFlag != 0 {
			s.SampleCompositionTimeOffset = int32(sr.Uint32())
		}

		b.Samples[i] = s
	}

	return box.FullRead(len(src))
}
