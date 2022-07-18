// Package sinf (Protection Scheme Information)
package sinf

import (
	"bytes"
	"fmt"

	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/box/scanner"
	"github.com/jwhittle933/streamline/media/mpeg/box/sinf/frma"
	"github.com/jwhittle933/streamline/media/mpeg/box/sinf/schi"
	"github.com/jwhittle933/streamline/media/mpeg/box/sinf/schm"
	"github.com/jwhittle933/streamline/media/mpeg/children"
)

const (
	SINF string = "sinf"
)

var (
	Children = children.Registry{
		frma.FRMA: frma.New,
		schi.SCHI: schi.New,
		schm.SCHM: schm.New,
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
	return SINF
}

func (b *Box) String() string {
	s :=  fmt.Sprintf("%s, children=%d", b.Info(), len(b.Children))

	for _, c := range b.Children {
		s += fmt.Sprintf("\n                %s", c.String())
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	s := scanner.New(bytes.NewReader(src))

	var err error
	if b.Children, err = s.ScanAllChildren(Children); err != nil {
		return 0, err
	}

	return box.FullRead(len(src))
}
