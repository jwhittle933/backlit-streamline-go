// Package mfro (Movie Fragment Random Access Offset)
package mfro

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/fullbox"
)

const (
	MFRO string = "mfro"
)

type Box struct {
	fullbox.Box
	ParentSize uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{*fullbox.New(i), 0}
}

func (Box) Type() string {
	return MFRO
}

func (b Box) String() string {
	return fmt.Sprintf("%s, parentsize=%d", b.Info(), b.ParentSize)
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)
	b.ParentSize = sr.Uint32()

	return box.FullRead(len(src))
}
