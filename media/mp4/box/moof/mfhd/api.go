// Package mfhd (Movie Fragment Header)
package mfhd

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/fullbox"
)

const (
	MFHD string = "mfhd"
)

type Box struct {
	fullbox.Box
	SequenceNumber uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{
		*fullbox.New(i),
		0,
	}
}

func (Box) Type() string {
	return MFHD
}

func (b Box) String() string {
	return fmt.Sprintf("%s, sequence=%d", b.Info(), b.SequenceNumber)
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)
	b.SequenceNumber = sr.Uint32()

	return box.FullRead(len(src))
}
