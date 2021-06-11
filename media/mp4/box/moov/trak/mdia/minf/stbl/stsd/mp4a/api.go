package mp4a

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
	"github.com/jwhittle933/streamline/media/mp4/box/children"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl/stsd/mp4a/esds"
	"github.com/jwhittle933/streamline/media/mp4/box/sample"
	"github.com/jwhittle933/streamline/media/mp4/box/scanner"
)

const (
	MP4A string = "mp4a"
)

var (
	Children = children.Registry{esds.ESDS: esds.New}
)

type Box struct {
	sample.Audio
	Children []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &Box{
		sample.Audio{
			Entry: sample.Entry{
				Box: base.Box{BoxInfo: i},
			},
		},
		make([]box.Boxed, 0),
	}
}

func (Box) Type() string {
	return MP4A
}

func (b *Box) String() string {
	s := fmt.Sprintf(
		"%s, boxes=%d, data_ref_index=%d status=\033[35mINCOMPLETE\033[0m",
		b.Info().String(),
		b.DataReferenceIndex,
		len(b.Children),
	)

	for _, c := range b.Children {
		s += fmt.Sprintf("\n              %s", c.String())
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	b.DataReferenceIndex = binary.BigEndian.Uint16(src[0:2])
	b.Version = binary.BigEndian.Uint16(src[2:4])
	// skip 6 for _reserved
	b.ChannelCount = binary.BigEndian.Uint16(src[10:12])

	s := scanner.New(bytes.NewReader(src[28:]))
	found, err := s.ScanAllChildren(Children)

	b.Children = found
	return len(src), err
}
