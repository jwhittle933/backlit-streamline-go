// Package mdia (Track Media Information)
package mdia

import (
	"bytes"
	"fmt"
	"io"

	"github.com/jwhittle933/streamline/media/mp4/base"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/hdlr"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/mdhd"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf"
	"github.com/jwhittle933/streamline/media/mp4/box/scanner"
	"github.com/jwhittle933/streamline/media/mp4/children"
)

const (
	MDIA string = "mdia"
)

var (
	Children = children.Registry{
		hdlr.HDLR: hdlr.New,
		mdhd.MDHD: mdhd.New,
		minf.MINF: minf.New,
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
	return MDIA
}

func (b Box) String() string {
	s := fmt.Sprintf("%s, boxes=%d", b.Info().String(), len(b.Children))

	for _, c := range b.Children {
		s += fmt.Sprintf("\n      %s", c.String())
	}

	return s
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	s := scanner.New(bytes.NewReader(src))

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
