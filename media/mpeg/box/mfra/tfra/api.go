// Package tfra (Track Fragment Random Access)
package tfra

import (
	"fmt"
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/fullbox"
)

const (
	TFRA string = "tfra"
)

type Box struct {
	fullbox.Box
	TrackID               uint32
	LengthSizeOfTrafNum   byte
	LengthSizeOfTrunNum   byte
	LengthSizeOfSampleNum byte
	NumberOfEntries       uint32
	Entries               []Entry
}

type Entry struct {
	Time        int64
	MoofOffset  int64
	TrafNumber  uint32
	TrunNumber  uint32
	SampleDelta uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{
		*fullbox.New(i),
		0,
		0,
		0,
		0,
		0,
		make([]Entry, 0, 0),
	}
}

func (Box) Type() string {
	return TFRA
}

func (b Box) String() string {
	s := fmt.Sprintf(
		"%s, version=%d, flags=%d, trackid=%d, ltrafnum=%d, ltrunnum=%d, lsamplenum=%d, entries=%d",
		b.Info(),
		b.Version,
		b.Flags,
		b.TrackID,
		b.LengthSizeOfTrafNum,
		b.LengthSizeOfTrunNum,
		b.LengthSizeOfSampleNum,
		b.NumberOfEntries,
	)

	for _, e := range b.Entries {
		s += fmt.Sprintf(
			"\n    [\033[1;34mentry\033[0m] time=%d, moofoffset=%d, trafnum=%d, trunnum=%d, delta=%d",
			e.Time,
			e.MoofOffset,
			e.TrafNumber,
			e.TrunNumber,
			e.SampleDelta,
		)
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)
	b.TrackID = sr.Uint32()
	sizeBlock := sr.Uint32()
	b.LengthSizeOfTrafNum = byte((sizeBlock >> 4) & 0x3)
	b.LengthSizeOfTrunNum = byte((sizeBlock >> 2) & 0x3)
	b.LengthSizeOfSampleNum = byte(sizeBlock & 0x3)

	b.NumberOfEntries = sr.Uint32()
	b.Entries = make([]Entry, b.NumberOfEntries)
	for i := 0; i < len(b.Entries); i++ {
		e := Entry{}

		if b.Version == 2 {
			e.Time = int64(sr.Uint64())
			e.MoofOffset = int64(sr.Uint64())
		} else {
			e.Time = int64(sr.Uint32())
			e.MoofOffset = int64(sr.Uint32())
		}

		switch b.LengthSizeOfTrafNum {
		case 0:
			e.TrafNumber = uint32(sr.Uint8())
		case 1:
			e.TrafNumber = uint32(sr.Uint16())
		case 2:
			e.TrafNumber = sr.Uint24()
		case 3:
			e.TrafNumber = sr.Uint32()
		}

		switch b.LengthSizeOfTrunNum {
		case 0:
			e.TrunNumber = uint32(sr.Uint8())
		case 1:
			e.TrunNumber = uint32(sr.Uint16())
		case 2:
			e.TrunNumber = sr.Uint24()
		case 3:
			e.TrunNumber = sr.Uint32()
		}

		switch b.LengthSizeOfSampleNum {
		case 0:
			e.SampleDelta = uint32(sr.Uint8())
		case 1:
			e.SampleDelta = uint32(sr.Uint16())
		case 2:
			e.SampleDelta = sr.Uint24()
		case 3:
			e.SampleDelta = sr.Uint32()
		}
	}

	return box.FullRead(len(src))
}
