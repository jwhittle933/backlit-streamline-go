// Package stbl (Sample Table)
package stbl

import (
	"bytes"
	"fmt"

	"github.com/jwhittle933/streamline/media/mp4/base"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl/co64"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl/ctts"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl/stco"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl/stsc"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl/stsd"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl/stsz"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl/stts"
	"github.com/jwhittle933/streamline/media/mp4/box/saio"
	"github.com/jwhittle933/streamline/media/mp4/box/saiz"
	"github.com/jwhittle933/streamline/media/mp4/box/sbgp"
	"github.com/jwhittle933/streamline/media/mp4/box/scanner"
	"github.com/jwhittle933/streamline/media/mp4/box/sgpd"
	"github.com/jwhittle933/streamline/media/mp4/box/stss"
	"github.com/jwhittle933/streamline/media/mp4/box/subs"
	"github.com/jwhittle933/streamline/media/mp4/children"
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
		stss.STSS: stss.New,
		co64.CO64: co64.New,
		ctts.CTTS: ctts.New,
		subs.SUBS: subs.New,
		saiz.SAIZ: saiz.New,
		saio.SAIO: saiz.New,
		sgpd.SGPD: sgpd.New,
		sbgp.SBGP: sbgp.New,
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
	s := fmt.Sprintf(
		"%s, children=%d",
		b.Info(),
		len(b.Children),
	)

	for _, c := range b.Children {
		s += fmt.Sprintf("\n          %s", c)
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	s := scanner.New(bytes.NewReader(src))
	found, err := s.ScanAllChildren(Children)
	b.Children = found

	return len(src), err
}
