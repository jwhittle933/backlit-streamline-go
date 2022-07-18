// Package stsc (Sample-to-Chunk)
package stsc

import (
	"encoding/binary"
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	STSC string = "stsc"
)

type Box struct {
	base.Box
	Version             uint8
	Flags               uint32
	EntryCount          uint32
	FirstChunk          []uint32
	SamplesPerChunk     []uint32
	SampleDescriptionID []uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		make([]uint32, 0),
		make([]uint32, 0),
		make([]uint32, 0),
	}
}

func (Box) Type() string {
	return STSC
}

func (b *Box) String() string {
	s := fmt.Sprintf(
		"%s, version=%d, flags=%d, samplecount=%d, samples=[",
		b.Info(),
		b.Version,
		b.Flags,
		b.EntryCount,
	)

	for i := uint32(0); i < b.EntryCount; i++ {
		s += fmt.Sprintf(
			"{chunk=%d, samples_per_chunk=%d, desc_id=%d}",
			b.FirstChunk[i],
			b.SamplesPerChunk[i],
			b.SampleDescriptionID[i],
		)
	}

	return s + "]"
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	versionAndFlags := sr.Uint32()

	b.Version = byte(versionAndFlags >> 24)
	b.Flags = versionAndFlags & box.FlagsMask
	b.EntryCount = sr.Uint32()

	b.FirstChunk = make([]uint32, b.EntryCount)
	b.SamplesPerChunk = make([]uint32, b.EntryCount)
	b.SampleDescriptionID = make([]uint32, b.EntryCount)
	for i := uint32(0); i < b.EntryCount; i++ {
		b.FirstChunk = append(b.FirstChunk, binary.BigEndian.Uint32(src[(8+12*i):(12+12*i)]))
		b.SamplesPerChunk = append(b.SamplesPerChunk, binary.BigEndian.Uint32(src[(12+12*i):(16+12*i)]))
		b.SampleDescriptionID = append(b.SampleDescriptionID, binary.BigEndian.Uint32(src[(16+12*i):(20+12*i)]))
	}

	return box.FullRead(len(src))
}

func lastItem(delta int, length int) string {
	if delta == length {
		return ""
	}

	return ", "
}
