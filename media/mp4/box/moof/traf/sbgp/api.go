// Package sbgp (Sample to Group)
package sbgp

import (
	"fmt"
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	SBGP string = "sbgp"
)

type Box struct {
	base.Box
	Version                 byte
	Flags                   uint32
	EntryCount              uint32
	GroupingType            string
	GroupingTypeParameter   uint32
	SampleCounts            []uint32
	GroupDescriptionIndices []uint32
}

type Entry struct {
	SampleCount           uint32
	GroupDescriptionIndex uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		"",
		0,
		make([]uint32, 0, 0),
		make([]uint32, 0, 0),
	}
}

func (b Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (Box) Type() string {
	return SBGP
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)

	versionAndFlags := sr.Uint32()
	b.Version = byte(versionAndFlags >> 24)
	b.Flags = versionAndFlags & box.FlagsMask
	b.GroupingType = sr.String(4)

	b.EntryCount = sr.Uint32()
	for i := uint32(0); i < b.EntryCount; i++ {
		b.SampleCounts = append(b.SampleCounts, sr.Uint32())
		b.GroupDescriptionIndices = append(b.GroupDescriptionIndices, sr.Uint32())
	}

	return box.FullRead(len(src))
}
