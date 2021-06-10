// Package mdia (Track Media Information)
package mdia

import (
	"bytes"
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	children2 "github.com/jwhittle933/streamline/media/mp4/box/children"
	hdlr2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/hdlr"
	mdhd2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/mdhd"
	minf2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf"
	scanner2 "github.com/jwhittle933/streamline/media/mp4/box/scanner"
	"io"
)

const (
	MDIA string = "mdia"
)

var (
	Children = children2.Registry{
		hdlr2.HDLR: hdlr2.New,
		mdhd2.MDHD: mdhd2.New,
		minf2.MINF: minf2.New,
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
	return MDIA
}

func (b Box) String() string {
	s := fmt.Sprintf("%s, boxes=%d", b.Info().String(), len(b.Children))

	for _, c := range b.Children {
		s += fmt.Sprintf("\n----->%s", c.String())
	}

	return s
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	s := scanner2.New(bytes.NewReader(src))

	for {
		child, err := s.ScanFor(Children)
		if err == io.EOF {
			return len(src), nil
		}

		if err != nil || child == nil {
			break
		}

		b.Children = append(b.Children, child)
	}

	return len(src), nil
}
