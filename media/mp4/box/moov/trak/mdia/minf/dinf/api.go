// Package dinf (Data Information)
package dinf

import (
	"bytes"
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	children2 "github.com/jwhittle933/streamline/media/mp4/box/children"
	dref2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/dinf/dref"
	scanner2 "github.com/jwhittle933/streamline/media/mp4/box/scanner"
)

const (
	DINF string = "dinf"
)

var (
	Children = children2.Registry{
		dref2.DREF: dref2.New,
	}
)

// Box is ISOBMFF dinf box type
type Box struct {
	base2.Box
	Children []box2.Boxed
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, make([]box2.Boxed, 0)}
}

func (Box) Type() string {
	return DINF
}

func (b *Box) String() string {
	s := fmt.Sprintf("%s, boxes=%d", b.Info().String(), len(b.Children))

	for _, c := range b.Children {
		s += fmt.Sprintf("\n--------->%s", c.String())
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	s := scanner2.New(bytes.NewReader(src))
	found, err := s.ScanAllChildren(Children)
	b.Children = found

	return len(src), err
}
