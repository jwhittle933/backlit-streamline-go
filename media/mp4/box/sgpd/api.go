// Package sgpd (Sample Group Description)
package sgpd

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
	"github.com/jwhittle933/streamline/media/mp4/box/children"
	"github.com/jwhittle933/streamline/media/mp4/box/sample/groupentry"
	"github.com/jwhittle933/streamline/media/mp4/box/sample/groupentry/alst"
	"github.com/jwhittle933/streamline/media/mp4/box/sample/groupentry/rap"
	"github.com/jwhittle933/streamline/media/mp4/box/sample/groupentry/roll"
	"github.com/jwhittle933/streamline/media/mp4/box/sample/groupentry/seig"
)

const (
	SGPD string = "sgpd"
)

var (
	GroupEntries = children.Registry{
		seig.SEIG: seig.New,
		roll.ROLL: roll.New,
		rap.RAP:   rap.New,
		alst.ALST: alst.New,
	}
)

// Box `sgpd` can appear in `stbl` or `traf`
type Box struct {
	base.Box
	Version                      byte
	Flags                        uint32
	GroupingType                 string
	DefaultLength                uint32
	DefaultGroupDescriptionIndex uint32
	EntryCount                   uint32
	DescriptionLengths           []uint32
	SampleGroupEntries           []groupentry.Sample
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		"",
		0,
		0,
		0,
		make([]uint32, 0, 0),
		make([]groupentry.Sample, 0, 0),
	}
}

func (b Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (Box) Type() string {
	return SGPD
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)

	versionAndFlags := sr.Uint32()
	b.Version = byte(versionAndFlags >> 24)
	b.Flags = versionAndFlags & box.FlagsMask
	b.GroupingType = sr.String(4)

	if b.Version >= 1 {
		b.DefaultLength = sr.Uint32()
	}

	if b.Version >= 2 {
		b.DefaultGroupDescriptionIndex = sr.Uint32()
	}

	b.EntryCount = sr.Uint32()
	b.DescriptionLengths = make([]uint32, b.EntryCount)
	for i := uint32(0); i < b.EntryCount; i++ {
		descriptionLength := sr.Uint32()

		if b.Version >= 1 && b.DefaultLength == 0 {
			descriptionLength = sr.Uint32()
			b.DescriptionLengths[i] = sr.Uint32()
		}

		entry, err := groupentry.ScanEntry(
			b.GroupingType,
			sr.Slice(int(descriptionLength)),
			GroupEntries.Get(b.GroupingType),
		)

		if err != nil {
			return 0, err
		}

		b.SampleGroupEntries = append(b.SampleGroupEntries, entry)
	}

	return box.FullRead(len(src))
}
