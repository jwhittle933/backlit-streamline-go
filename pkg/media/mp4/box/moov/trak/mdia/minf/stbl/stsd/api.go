// Package stsd (Sample Descriptions)
package stsd

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/children"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsd/avc1"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsd/avcC"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsd/mp4a"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsd/pasp"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/scanner"
)

const (
	STSD string = "stsd"
)

var (
	Children = children.Registry{
		avc1.AVC1: avc1.New,
		avcC.AVCC: avcC.New,
		mp4a.MP4A: mp4a.New,
		pasp.PASP: pasp.New,
	}
)

type Box struct {
	base.Box
	Version    uint8
	Flags      uint32
	EntryCount uint32
	Entries    Entries
	Children   []box.Boxed
	raw        []byte
}

type Entry struct {
	DataRefIndex uint16
	Sample       []byte
}

type Entries []Entry

func (e Entries) String() string {
	s := ""
	for _, entry := range e {
		s += fmt.Sprintf("{data_ref_index=%d, sample=%s} ", entry.DataRefIndex, entry.Sample)
	}

	return s
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		make([]Entry, 0),
		make([]box.Boxed, 0),
		make([]byte, 0),
	}
}

func (Box) Type() string {
	return STSD
}

func (b *Box) String() string {
	s := b.Info().String() + fmt.Sprintf(
		", version=%d, flags=%d, entry_count=%d, entries=[%s] status=\033[35mINCOMPLETE\033[0m",
		b.Version,
		b.Flags,
		b.EntryCount,
		b.Entries.String(),
	)

	for _, c := range b.Children {
		s += fmt.Sprintf("\n----------->%s", c.String())
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	b.Version = src[0]
	b.Flags = binary.BigEndian.Uint32([]byte{0x00, src[1], src[2], src[3]})
	b.raw = src[4:]

	b.EntryCount = binary.BigEndian.Uint32(b.raw[0:4])

	s := scanner.New(bytes.NewReader(b.raw[4:]))
	found, err := s.ScanAllChildren(Children)
	b.Children = found

	return len(src), err
}
