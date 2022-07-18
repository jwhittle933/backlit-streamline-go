// Package mvex (Movie Extends)
package mvex

import (
	"bytes"
	"fmt"

	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/box/moov/mvex/mehd"
	"github.com/jwhittle933/streamline/media/mpeg/box/moov/mvex/trex"
	"github.com/jwhittle933/streamline/media/mpeg/box/scanner"
	"github.com/jwhittle933/streamline/media/mpeg/children"
)

const (
	MVEX string = "mvex"
)

var (
	Children = children.Registry{
		mehd.MEHD: mehd.New,
		trex.TREX: trex.New,
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
	return MVEX
}

func (b Box) String() string {
	s := fmt.Sprintf(
		"%s, children=%d",
		b.Info(),
		len(b.Children),
	)

	for _, c := range b.Children {
		s += fmt.Sprintf("\n    %s", c)
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
