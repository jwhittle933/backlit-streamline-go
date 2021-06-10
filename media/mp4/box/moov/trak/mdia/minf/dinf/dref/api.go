// Package dref (Data Reference), declare source of media data in track
package dref

import (
	"bytes"
	"encoding/binary"
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	children2 "github.com/jwhittle933/streamline/media/mp4/box/children"
	url2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/dinf/dref/url"
	urn2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/dinf/dref/urn"
	scanner2 "github.com/jwhittle933/streamline/media/mp4/box/scanner"
)

const (
	DREF   string = "dref"
	inFile        = 0x01
)

var (
	Children = children2.Registry{
		url2.URL: url2.New,
		urn2.URN: urn2.New,
	}
)

// Box is ISOBMFF dref box type
type Box struct {
	base2.Box
	Version    uint8
	Flags      uint32
	EntryCount uint32
	Children   []box2.Boxed
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, 0, 0, 0, make([]box2.Boxed, 0)}
}

func (Box) Type() string {
	return DREF
}

func (b *Box) String() string {
	s := b.Info().String() + fmt.Sprintf(
		", version=%d, flags=%d, entry_count=%d",
		b.Version,
		b.Flags,
		b.EntryCount,
	)

	for _, c := range b.Children {
		s += fmt.Sprintf("\n----------->%s", c.String())
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	b.Version = src[0]
	b.Flags = binary.BigEndian.Uint32([]byte{0x00, src[1], src[2], src[3]})
	b.EntryCount = binary.BigEndian.Uint32(src[4:8])

	s := scanner2.New(bytes.NewReader(src[8:]))
	found, err := s.ScanAllChildren(Children)
	b.Children = found

	return len(src), err
}
