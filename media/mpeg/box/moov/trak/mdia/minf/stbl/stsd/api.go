// Package stsd (Sample Descriptions)
package stsd

import (
	"bytes"
	"fmt"
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/box/sample/audio"
	"github.com/jwhittle933/streamline/media/mpeg/box/sample/visual"
	"github.com/jwhittle933/streamline/media/mpeg/box/scanner"
	"github.com/jwhittle933/streamline/media/mpeg/children"
	"github.com/jwhittle933/streamline/media/mpeg/fullbox"
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
		"encv": visual.New,
		//"c608": audio.New,
	}
)

type Box struct {
	fullbox.Box
	SampleCount uint32
	Children    []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &Box{
		*fullbox.New(i),
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
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)
	b.SampleCount = sr.Uint32()

	s := scanner.New(bytes.NewReader(sr.Remaining()))
	var err error
	b.Children, err = s.ScanAllChildren(Children)
	if err != nil {
		return 0, err
	}

	return box.FullRead(len(src))
}
