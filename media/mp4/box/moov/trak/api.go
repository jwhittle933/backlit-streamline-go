// Package trak (Track)
// See: https://github.com/itsjamie/bmff/blob/master/trak.go
package trak

import (
	"bytes"
	"fmt"
	udta2 "github.com/jwhittle933/streamline/media/mp4/box/udta"

	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
	"github.com/jwhittle933/streamline/media/mp4/box/children"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/edts"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/tkhd"
	"github.com/jwhittle933/streamline/media/mp4/box/scanner"
)

const (
	TRAK string = "trak"
)

var Children = children.Registry{
	tkhd.TKHD:  tkhd.New,
	edts.EDTS:  edts.New,
	mdia.MDIA:  mdia.New,
	udta2.UDTA: udta2.New,
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
