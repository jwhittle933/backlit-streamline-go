// Pacakge elst (Edit List Box)
package elst

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/jwhittle933/streamline/media/mp4/base"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
)

const (
	ELST string = "elst"
)

var (
	InvalidELST = errors.New("box elst must have entry_count of 1 or be excluded (Common File Format & Media Formats Specification Version 2.1, section 2.1.2.2")
)

// Box is ISOBMFF elst box type
type Box struct {
	base.Box
	Version    uint8
	Flags      uint32
	EntryCount uint32
	Entries    []Entry
	raw        []byte
}

type Entry struct {
	SegmentDuration   uint64
	MediaTime         int64
	MediaRateInteger  int16
	MediaRateFraction int16
}

func New(i *box2.Info) box2.Boxed {
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
	return ELST
}

func (b Box) String() string {
	return b.Info().String() + fmt.Sprintf(
		", version=%d, flags=%d, entry_count=%d, entries=[media_rate_fraction=%d, media_rate=%d, media_time=%d, segment_duration=%d]",
		b.Version,
		b.Flags,
		b.EntryCount,
		b.Entries[0].MediaRateFraction,
		b.Entries[0].MediaRateInteger,
		b.Entries[0].MediaTime,
		b.Entries[0].SegmentDuration,
	)
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	b.Version = src[0]
	b.Flags = binary.BigEndian.Uint32([]byte{0x00, src[1], src[2], src[3]})

	b.raw = src[4:]
	b.EntryCount = binary.BigEndian.Uint32(b.raw[0:4])

	if b.EntryCount != 1 {
		return 8, InvalidELST
	}

	offset := 4
	if b.Version == 1 {
		entry := Entry{
			SegmentDuration: binary.BigEndian.Uint64(b.raw[offset : offset+8]),
			MediaTime:       int64(binary.BigEndian.Uint64(b.raw[offset+8 : offset+16])),
		}
		offset += 16

		entry.MediaRateInteger = int16(binary.BigEndian.Uint16(b.raw[offset : offset+2]))
		entry.MediaRateFraction = int16(binary.BigEndian.Uint16(b.raw[offset+2 : offset+4]))
		offset += 4

		b.Entries = []Entry{entry}
	} else if b.Version == 0 {
		entry := Entry{
			SegmentDuration: uint64(binary.BigEndian.Uint32(b.raw[offset : offset+4])),
			MediaTime:       int64(binary.BigEndian.Uint32(b.raw[offset+4 : offset+8])),
		}
		offset += 8

		entry.MediaRateInteger = int16(binary.BigEndian.Uint16(b.raw[offset : offset+2]))
		entry.MediaRateFraction = int16(binary.BigEndian.Uint16(b.raw[offset+2 : offset+4]))
		offset += 4

		b.Entries = []Entry{entry}
	}

	return len(src), nil
}
