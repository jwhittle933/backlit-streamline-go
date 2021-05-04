// Package stbl (Sample Table)
package stbl

import (
	"bytes"
	"fmt"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/children"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stco"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsc"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsd"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsz"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stts"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/scanner"
)

const (
	STBL string = "stbl"
)

var (
	Children = children.Registry{
		stco.STCO: stco.New,
		stsc.STSC: stsc.New,
		stsd.STSD: stsd.New,
		stsz.STSZ: stsz.New,
		stts.STTS: stts.New,
	}
)

type Box struct {
	base.Box
	Children []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, make([]box.Boxed, 0)}
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
	s := scanner.New(bytes.NewReader(src))
	found, err := s.ScanAllChildren(Children)
	b.Children = found

	return len(src), err
}
