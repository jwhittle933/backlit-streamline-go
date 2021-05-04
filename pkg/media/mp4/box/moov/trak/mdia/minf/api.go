// Package minf (Media Information)
package minf

import (
	"bytes"
	"fmt"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/children"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/dinf"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/smhd"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/sthd"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/vmhd"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/scanner"
)

const (
	MINF string = "minf"
)

var (
	Children = children.Registry{
		dinf.DINF: dinf.New,
		smhd.SMHD: smhd.New,
		stbl.STBL: stbl.New,
		sthd.STHD: sthd.New,
		vmhd.VMHD: vmhd.New,
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
	return MINF
}

func (b *Box) String() string {
	s := fmt.Sprintf("%s, boxes=%d", b.Info().String(), len(b.Children))

	for _, c := range b.Children {
		s += fmt.Sprintf("\n------->%s", c.String())
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	s := scanner.New(bytes.NewReader(src))
	found, err := s.ScanAllChildren(Children)
	b.Children = found

	return len(src), err
}
