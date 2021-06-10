// Package trak (Track)
// See: https://github.com/itsjamie/bmff/blob/master/trak.go
package trak

import (
	"bytes"
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	children2 "github.com/jwhittle933/streamline/media/mp4/box/children"
	edts2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/edts"
	mdia2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia"
	tkhd2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/tkhd"
	udta2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/udta"
	scanner2 "github.com/jwhittle933/streamline/media/mp4/box/scanner"
)

const (
	TRAK string = "trak"
)

var Children = children2.Registry{
	tkhd2.TKHD: tkhd2.New,
	edts2.EDTS: edts2.New,
	mdia2.MDIA: mdia2.New,
	udta2.UDTA: udta2.New,
}

type Box struct {
	base2.Box
	Children []box2.Boxed
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, make([]box2.Boxed, 0)}
}

func (Box) Type() string {
	return TRAK
}

func (b Box) String() string {
	s := fmt.Sprintf("%s, boxes=%d", b.Info(), len(b.Children))

	for _, c := range b.Children {
		s += fmt.Sprintf("\n--->%s", c.String())
	}

	return s
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	s := scanner2.New(bytes.NewReader(src))

	var err error
	b.Children, err = s.ScanAllChildren(Children)
	if err != nil {
		return 0, err
	}

	return len(src), nil
}
