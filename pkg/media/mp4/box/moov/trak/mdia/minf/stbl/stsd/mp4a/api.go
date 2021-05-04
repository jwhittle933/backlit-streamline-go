package mp4a

import (
	"bytes"
	"fmt"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/children"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsd/mp4a/esds"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/scanner"
)

const (
	MP4A string = "mp4a"
)

var (
	Children = children.Registry{
		esds.ESDS: esds.New,
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
	return MP4A
}

func (b *Box) String() string {
	s := fmt.Sprintf("%s, boxes=%d, status=\033[35mINCOMPLETE\033[0m", b.Info().String(), len(b.Children))

	for _, c := range b.Children {
		s += fmt.Sprintf("\n------------->%s", c.String())
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	s := scanner.New(bytes.NewReader(src[28:]))
	found, err := s.ScanAllChildren(Children)

	b.Children = found
	return len(src), err
}
