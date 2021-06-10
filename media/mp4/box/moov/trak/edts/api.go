// Package edts (Edit Box)
// See: https://github.com/itsjamie/bmff/blob/b0eabe94928cc481a1b373fc5e920257ab464912/edts.go
package edts

import (
	"bytes"
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	children2 "github.com/jwhittle933/streamline/media/mp4/box/children"
	elst2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/edts/elst"
	scanner2 "github.com/jwhittle933/streamline/media/mp4/box/scanner"
	"io"
)

const (
	EDTS string = "edts"
)

var Children = children2.Registry{elst2.ELST: elst2.New}

// Box is ISOBMFF edts box type
type Box struct {
	base2.Box
	Children []box2.Boxed
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, make([]box2.Boxed, 0)}
}

func (Box) Type() string {
	return EDTS
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
