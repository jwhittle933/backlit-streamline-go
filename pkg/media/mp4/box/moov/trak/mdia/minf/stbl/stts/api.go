// Package stts (Time to Sample) decoding
package stts

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	STTS string = "stts"
)

type Box struct {
	base.Box
	EntryCount uint32
	Entries    []Entry
}

type Entry struct {
	SampleCount uint32
	SampleDelta uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, 0, make([]Entry, 0)}
}

func (Box) Type() string {
	return STTS
}

func (b *Box) String() string {
	return b.Info().String()
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
