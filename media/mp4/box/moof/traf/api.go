// Package traf (Track Fragment)
package traf

import (
	"bytes"
	"fmt"

	"github.com/jwhittle933/streamline/media/mp4/base"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/moof/traf/tfdt"
	"github.com/jwhittle933/streamline/media/mp4/box/moof/traf/tfhd"
	"github.com/jwhittle933/streamline/media/mp4/box/moof/traf/trun"
	"github.com/jwhittle933/streamline/media/mp4/box/scanner"
	"github.com/jwhittle933/streamline/media/mp4/box/subs"
	"github.com/jwhittle933/streamline/media/mp4/children"
)

const (
	TRAF string = "traf"
)

var (
	Children = children.Registry{
		subs.SUBS: subs.New,
		tfhd.TFHD: tfhd.New,
		tfdt.TFDT: tfdt.New,
		trun.TRUN: trun.New,
	}
)

type Box struct {
	base.Box
	Children []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, make([]box.Boxed, 0)}
}

func (b Box) String() string {
	s := fmt.Sprintf("%s", b.Info())

	for _, c := range b.Children {
		s += fmt.Sprintf("\n    %s", c)
	}

	return s
}

func (Box) Type() string {
	return TRAF
}

func (b *Box) Write(src []byte) (int, error) {
	s := scanner.New(bytes.NewReader(src))

	var err error
	b.Children, err = s.ScanAllChildren(Children)
	if err != nil {
		return 0, err
	}

	return box.FullRead(len(src))
}
