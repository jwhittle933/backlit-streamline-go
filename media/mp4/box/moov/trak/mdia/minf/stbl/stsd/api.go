// Package stsd (Sample Descriptions)
package stsd

import (
	"bytes"
	"encoding/binary"
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	children2 "github.com/jwhittle933/streamline/media/mp4/box/children"
	mp4a2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl/stsd/mp4a"
	visual2 "github.com/jwhittle933/streamline/media/mp4/box/sample/visual"
	scanner2 "github.com/jwhittle933/streamline/media/mp4/box/scanner"
)

const (
	STSD string = "stsd"
)

var (
	Children = children2.Registry{
		"avc1":     visual2.New,
		"avc3":     visual2.New,
		"hev1":     visual2.New,
		"hvc1":     visual2.New,
		mp4a2.MP4A: mp4a2.New,
	}
)

type Box struct {
	base2.Box
	Version     uint8
	Flags       uint32
	SampleCount uint32
	Children    []box2.Boxed
}

func New(i *box2.Info) box2.Boxed {
	return &Box{
		base2.Box{BoxInfo: i},
		0,
		0,
		0,
		make([]box2.Boxed, 0),
	}
}

func (Box) Type() string {
	return STSD
}

func (b *Box) String() string {
	s := b.Info().String() + fmt.Sprintf(
		", version=%d, flags=%d, sample_count=%d",
		b.Version,
		b.Flags,
		b.SampleCount,
	)

	for _, c := range b.Children {
		s += fmt.Sprintf("\n----------->%s", c.String())
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	b.Version = src[0]
	b.Flags = binary.BigEndian.Uint32([]byte{0x00, src[1], src[2], src[3]})
	b.SampleCount = binary.BigEndian.Uint32(src[4:8])

	s := scanner2.New(bytes.NewReader(src[8:]))
	found, err := s.ScanAllChildren(Children)
	b.Children = found

	return len(src), err
}
