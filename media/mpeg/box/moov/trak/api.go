// Package trak (Track)
// See: https://github.com/itsjamie/bmff/blob/master/trak.go
package trak

import (
	"bytes"
	"fmt"

	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/box/moov/trak/edts"
	"github.com/jwhittle933/streamline/media/mpeg/box/moov/trak/mdia"
	"github.com/jwhittle933/streamline/media/mpeg/box/moov/trak/tkhd"
	"github.com/jwhittle933/streamline/media/mpeg/box/scanner"
	"github.com/jwhittle933/streamline/media/mpeg/box/udta"
	"github.com/jwhittle933/streamline/media/mpeg/children"
)

const (
	TRAK string = "trak"
)

var Children = children.Registry{
	tkhd.TKHD: tkhd.New,
	edts.EDTS: edts.New,
	mdia.MDIA: mdia.New,
	udta.UDTA: udta.New,
}

type Box struct {
	base.Box
	Children []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, make([]box.Boxed, 0)}
}

func (Box) Type() string {
	return TRAK
}

func (b Box) String() string {
	s := fmt.Sprintf("%s, boxes=%d", b.Info(), len(b.Children))

	for _, c := range b.Children {
		s += fmt.Sprintf("\n    %s", c.String())
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

	return len(src), nil
}
