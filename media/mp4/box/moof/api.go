// Package moof (Movie Fragment)
package moof

import (
	"bytes"
	"fmt"

	"github.com/jwhittle933/streamline/media/mp4/base"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/moof/mfhd"
	"github.com/jwhittle933/streamline/media/mp4/box/moof/traf"
	"github.com/jwhittle933/streamline/media/mp4/box/scanner"
	"github.com/jwhittle933/streamline/media/mp4/children"
)

const (
	MOOF string = "moof"
)

var (
	Children = children.Registry{
		mfhd.MFHD: mfhd.New,
		traf.TRAF: traf.New,
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
	return MOOF
}

func (b Box) String() string {
	s := fmt.Sprintf(
		"%s",
		b.Info(),
	)

	for _, c := range b.Children {
		s += fmt.Sprintf("\n  %s", c)
	}

	return s
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	s := scanner.New(bytes.NewReader(src))

	var err error
	b.Children, err = s.ScanAllChildren(Children)
	if err != nil {
		return 0, err
	}

	return box.FullRead(len(src))
}
