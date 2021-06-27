// Package dref (Data Reference), declare source of media data in track
package dref

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/jwhittle933/streamline/media/mp4/children"

	"github.com/jwhittle933/streamline/media/mp4/base"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/dinf/dref/url"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak/mdia/minf/dinf/dref/urn"
	"github.com/jwhittle933/streamline/media/mp4/box/scanner"
)

const (
	DREF   string = "dref"
	inFile        = 0x01
)

var (
	Children = children.Registry{
		url.URL: url.New,
		urn.URN: urn.New,
	}
)

// Box is ISOBMFF dref box type
type Box struct {
	base.Box
	Version    uint8
	Flags      uint32
	EntryCount uint32
	Children   []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		make([]box.Boxed, 0),
	}
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
		s += fmt.Sprintf("\n            %s", c.String())
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	b.Version = src[0]
	b.Flags = binary.BigEndian.Uint32([]byte{0x00, src[1], src[2], src[3]})
	b.EntryCount = binary.BigEndian.Uint32(src[4:8])

	s := scanner.New(bytes.NewReader(src[8:]))
	found, err := s.ScanAllChildren(Children)
	b.Children = found

	return len(src), err
}
