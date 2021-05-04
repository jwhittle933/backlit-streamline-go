// Package sinf (Protection Scheme Information)
package sinf

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/children"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsc/sinf/frma"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsc/sinf/schi"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsc/sinf/schm"
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
