// Package sinf (Protection Scheme Information)
package sinf

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/children"
	frma2 "github.com/jwhittle933/streamline/pkg/media/mp4/box/sinf/frma"
	schi2 "github.com/jwhittle933/streamline/pkg/media/mp4/box/sinf/schi"
	schm2 "github.com/jwhittle933/streamline/pkg/media/mp4/box/sinf/schm"
)

const (
	SINF string = "sinf"
)

var (
	Children = children.Registry{
		frma2.FRMA: frma2.New,
		schi2.SCHI: schi2.New,
		schm2.SCHM: schm2.New,
	}
)

type Box struct {
	base.Box
}

func New(i *box.Info) box.Boxed {
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
