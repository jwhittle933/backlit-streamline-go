// Package stts (Time to Sample) decoding
package stts

import (
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	STTS string = "stts"
)

type Box struct {
	base2.Box
	EntryCount uint32
	Entries    []Entry
}

type Entry struct {
	SampleCount uint32
	SampleDelta uint32
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, 0, make([]Entry, 0)}
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
