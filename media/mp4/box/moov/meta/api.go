// Package meta for multi-track required meta-data
package meta

import (
	"bytes"
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/ilst"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/hdlr"
	"github.com/jwhittle933/streamline/media/mp4/box/scanner"
	"github.com/jwhittle933/streamline/media/mp4/children"
	"github.com/jwhittle933/streamline/media/mp4/fullbox"
)

const (
	META string = "meta"
)

var (
	Children = children.Registry{
		hdlr.HDLR: hdlr.New,
		ilst.ILST: ilst.New,
	}
)

// Box is ISO BMFF meta box type
type Box struct {
	fullbox.Box
	Children []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &Box{*fullbox.New(i), make([]box.Boxed, 0)}
}

func (Box) Type() string {
	return META
}

func (b *Box) String() string {
	s := fmt.Sprintf(
		"%s",
		b.Info(),
	)

	for _, c := range b.Children {
		s += fmt.Sprintf("\n      %s", c)
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)

	s := scanner.New(bytes.NewReader(sr.Remaining()))

	var err error
	b.Children, err = s.ScanAllChildren(Children)
	if err != nil {
		return 0, err
	}

	return box.FullRead(len(src))
}
