// Package sinf (Protection Scheme Information)
package sinf

import (
	"github.com/jwhittle933/streamline/media/mp4/base"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/sinf/frma"
	"github.com/jwhittle933/streamline/media/mp4/box/sinf/schi"
	"github.com/jwhittle933/streamline/media/mp4/box/sinf/schm"
	"github.com/jwhittle933/streamline/media/mp4/children"
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
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base.Box{BoxInfo: i}}
}

func (Box) Type() string {
	return SINF
}

func (b *Box) String() string {
	return b.Info().String()
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
