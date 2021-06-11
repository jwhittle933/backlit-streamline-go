// Package stsd (Sample Descriptions)
package stsd

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
	"github.com/jwhittle933/streamline/media/mp4/box/children"
	"github.com/jwhittle933/streamline/media/mp4/box/sample/audio"
	"github.com/jwhittle933/streamline/media/mp4/box/sample/visual"
	"github.com/jwhittle933/streamline/media/mp4/box/scanner"
)

const (
	STSD string = "stsd"
)

var (
	Children = children.Registry{
		"avc1": visual.New,
		"avc3": visual.New,
		"hev1": visual.New,
		"hvc1": visual.New,
		"mp4a": audio.New,
	}
)

type Box struct {
	base.Box
	Version     uint8
	Flags       uint32
	SampleCount uint32
	Children    []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		make([]box.Boxed, 0),
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
		s += fmt.Sprintf("\n            %s", c.String())
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	b.Version = src[0]
	b.Flags = binary.BigEndian.Uint32([]byte{0x00, src[1], src[2], src[3]})
	b.SampleCount = binary.BigEndian.Uint32(src[4:8])

	s := scanner.New(bytes.NewReader(src[8:]))
	found, err := s.ScanAllChildren(Children)
	b.Children = found

	return len(src), err
}
