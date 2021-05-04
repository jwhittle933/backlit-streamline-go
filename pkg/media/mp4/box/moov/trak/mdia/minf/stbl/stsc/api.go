// Package stsc (Sample-to-Chunk)
package stsc

import (
	"bytes"
	"fmt"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/children"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsc/sinf"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/scanner"
)

const (
	STSC string = "stsc"
)

var (
	Children = children.Registry{
		sinf.SINF: sinf.New,
	}
)

type Box struct {
	base.Box
	EntryCount uint32
	Entries    []Entry
	Children   []box.Boxed
}

type Entry struct {
	FirstChunk             uint32
	SamplesPerChunk        uint32
	SampleDescriptionIndex uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, 0, make([]Entry, 0), make([]box.Boxed, 0)}
}

func (Box) Type() string {
	return STSC
}

func (b *Box) String() string {
	s := fmt.Sprintf("%s, boxes=%d", b.Info().String(), len(b.Children))

	for _, c := range b.Children {
		s += fmt.Sprintf("\n----------->%s", c.String())
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	s := scanner.New(bytes.NewReader(src))
	found, err := s.ScanAllChildren(Children)
	b.Children = found

	return len(src), err
}
