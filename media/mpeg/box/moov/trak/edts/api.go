// Package edts (Edit Box)
// See: https://github.com/itsjamie/bmff/blob/b0eabe94928cc481a1b373fc5e920257ab464912/edts.go
package edts

import (
	"bytes"
	"fmt"
	"io"

	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/box/moov/trak/edts/elst"
	"github.com/jwhittle933/streamline/media/mpeg/box/scanner"
	"github.com/jwhittle933/streamline/media/mpeg/children"
)

const (
	EDTS string = "edts"
)

var Children = children.Registry{elst.ELST: elst.New}

// Box is ISOBMFF edts box type
type Box struct {
	base.Box
	Children []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, make([]box.Boxed, 0)}
}

func (Box) Type() string {
	return EDTS
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
