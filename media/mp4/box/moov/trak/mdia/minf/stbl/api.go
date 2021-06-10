// Package stbl (Sample Table)
package stbl

import (
	"bytes"
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	children2 "github.com/jwhittle933/streamline/media/mp4/box/children"
	stco2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl/stco"
	stsc2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl/stsc"
	stsd2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl/stsd"
	stsz2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl/stsz"
	stts2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl/stts"
	scanner2 "github.com/jwhittle933/streamline/media/mp4/box/scanner"
)

const (
	STBL string = "stbl"
)

var (
	Children = children2.Registry{
		stco2.STCO: stco2.New,
		stsc2.STSC: stsc2.New,
		stsd2.STSD: stsd2.New,
		stsz2.STSZ: stsz2.New,
		stts2.STTS: stts2.New,
	}
)

type Box struct {
	base2.Box
	Children []box2.Boxed
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, make([]box2.Boxed, 0)}
}

func (Box) Type() string {
	return STBL
}

func (b *Box) String() string {
	s := fmt.Sprintf("%s, boxes=%d, status=\033[35mINCOMPLETE\033[0m", b.Info().String(), len(b.Children))

	for _, c := range b.Children {
		s += fmt.Sprintf("\n--------->%s", c.String())
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	s := scanner2.New(bytes.NewReader(src))
	found, err := s.ScanAllChildren(Children)
	b.Children = found

	return len(src), err
}
