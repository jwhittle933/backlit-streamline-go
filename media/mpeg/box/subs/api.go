package subs

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/fullbox"
)

const (
	SUBS = "subs"
)

type Box struct {
	fullbox.Box
	Entries []Entry
}

type Entry struct {
	SampleData uint32
	Samples    []Sample
}

type Sample struct {
	Size                uint32
	CodecSpecificParams uint32
	Priority            uint8
	Discardable         uint8
}

func New(i *box.Info) box.Boxed {
	return &Box{
		*fullbox.New(i),
		make([]Entry, 0, 0),
	}
}

func (Box) Type() string {
	return SUBS
}

func (b Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)

	b.Entries = make([]Entry, sr.Uint32())
	for i := 0; i < len(b.Entries); i++ {
		e := Entry{SampleData: sr.Uint32(), Samples: make([]Sample, sr.Uint16())}

		for j := 0; j < len(e.Samples); j++ {
			s := Sample{}
			if b.Version == 1 {
				s.Size = sr.Uint32()
			} else {
				s.Size = uint32(sr.Uint16())
			}

			s.Priority = sr.Uint8()
			s.Discardable = sr.Uint8()
			s.CodecSpecificParams = sr.Uint32()
			e.Samples[j] = s
		}

		b.Entries[i] = e
	}

	return box.FullRead(len(src))
}
