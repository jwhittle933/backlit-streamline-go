package mp4a

import (
	"bytes"
	"encoding/binary"
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	children2 "github.com/jwhittle933/streamline/media/mp4/box/children"
	esds2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/stbl/stsd/mp4a/esds"
	sample2 "github.com/jwhittle933/streamline/media/mp4/box/sample"
	scanner2 "github.com/jwhittle933/streamline/media/mp4/box/scanner"
)

const (
	MP4A string = "mp4a"
)

var (
	Children = children2.Registry{
		esds2.ESDS: esds2.New,
	}
)

type Box struct {
	sample2.Audio
	Children []box2.Boxed
}

func New(i *box2.Info) box2.Boxed {
	return &Box{
		sample2.Audio{
			Entry: sample2.Entry{
				Box: base2.Box{BoxInfo: i},
			},
		},
		make([]box2.Boxed, 0),
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
		s += fmt.Sprintf("\n------------->%s", c.String())
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	b.DataReferenceIndex = binary.BigEndian.Uint16(src[0:2])
	b.Version = binary.BigEndian.Uint16(src[2:4])
	// skip 6 for _reserved
	b.ChannelCount = binary.BigEndian.Uint16(src[10:12])

	s := scanner2.New(bytes.NewReader(src[28:]))
	found, err := s.ScanAllChildren(Children)

	b.Children = found
	return len(src), err
}
