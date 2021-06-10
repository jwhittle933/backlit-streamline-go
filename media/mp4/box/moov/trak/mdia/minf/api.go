// Package minf (Media Information)
package minf

import (
	"bytes"
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	children2 "github.com/jwhittle933/streamline/media/mp4/box/children"
	dinf2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/dinf"
	smhd2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/smhd"
	stbl2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl"
	sthd2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/sthd"
	vmhd2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/vmhd"
	scanner2 "github.com/jwhittle933/streamline/media/mp4/box/scanner"
)

const (
	MINF string = "minf"
)

var (
	Children = children2.Registry{
		dinf2.DINF: dinf2.New,
		smhd2.SMHD: smhd2.New,
		stbl2.STBL: stbl2.New,
		sthd2.STHD: sthd2.New,
		vmhd2.VMHD: vmhd2.New,
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
	s := scanner2.New(bytes.NewReader(src))

	var err error
	b.Children, err = s.ScanAllChildren(Children)
	if err != nil {
		return 0, err
	}

	return len(src), err
}
