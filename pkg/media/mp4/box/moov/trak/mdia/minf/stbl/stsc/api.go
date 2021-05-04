// Package stsc (Sample-to-Chunk)
package stsc

import (
	"encoding/binary"
	"fmt"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	STSC string = "stsc"
)

type Box struct {
	base.Box
	Version    uint8
	Flags      uint32
	EntryCount uint32
	Entries    Entries
	raw        []byte
}

type Entry struct {
	FirstChunk             uint32
	SamplesPerChunk        uint32
	SampleDescriptionIndex uint32
}

type Entries []Entry

func (e Entries) String() string {
	if len(e) > 2 {
		return shortString(e)
	}

	return verboseString(e)
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		make([]Entry, 0),
		make([]byte, 0),
	}
}

func (Box) Type() string {
	return STSC
}

func (b *Box) String() string {
	s := fmt.Sprintf(
		"%s, version=%d, flags=%d, entry_count=%d, entries=[%s]",
		b.Info().String(),
		b.Version,
		b.Flags,
		b.EntryCount,
		b.Entries.String(),
	)

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	b.Version = src[0]
	b.Flags = binary.BigEndian.Uint32([]byte{0x00, src[1], src[2], src[3]})
	b.raw = src[4:]

	b.EntryCount = binary.BigEndian.Uint32(b.raw[0:4])
	offset := 4
	for i := 0; uint32(i) < b.EntryCount; i++ {
		entry := Entry{
			FirstChunk:             binary.BigEndian.Uint32(b.raw[offset : offset+4]),
			SamplesPerChunk:        binary.BigEndian.Uint32(b.raw[offset+4 : offset+8]),
			SampleDescriptionIndex: binary.BigEndian.Uint32(b.raw[offset+8 : offset+12]),
		}

		offset += 12
		b.Entries = append(b.Entries, entry)
	}

	return len(src), nil
}
func shortString(e Entries) string {
	var s string

	for i, entry := range e {
		s += fmt.Sprintf(
			"{%d, %d, %d}%s",
			entry.FirstChunk,
			entry.SamplesPerChunk,
			entry.SampleDescriptionIndex,
			lastItem(i, len(e)-1),
		)
	}

	return s
}

func verboseString(e Entries) string {
	var s string

	for i, entry := range e {
		s += fmt.Sprintf(
			"{first_chunk=%d, samples_per_chunk=%d, sample_description_index=%d}%s",
			entry.FirstChunk,
			entry.SamplesPerChunk,
			entry.SampleDescriptionIndex,
			lastItem(i, len(e)-1),
		)
	}

	return s
}

func lastItem(delta int, length int) string {
	if delta == length {
		return ""
	}

	return ", "
}
