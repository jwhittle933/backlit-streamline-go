// Package sinf (Protection Scheme Information)
package sinf

import (
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	children2 "github.com/jwhittle933/streamline/media/mp4/box/children"
	"github.com/jwhittle933/streamline/media/mp4/box/sinf/frma"
	"github.com/jwhittle933/streamline/media/mp4/box/sinf/schi"
	"github.com/jwhittle933/streamline/media/mp4/box/sinf/schm"
)

const (
	SINF string = "sinf"
)

var (
	Children = children2.Registry{
		frma.FRMA: frma.New,
		schi.SCHI: schi.New,
		schm.SCHM: schm.New,
	}
)

type Box struct {
	base2.Box
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}}
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
