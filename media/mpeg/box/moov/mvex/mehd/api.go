// Package mehd (Movie Extends)
package mehd

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/fullbox"
)

const (
	MEHD string = "mehd"
)

type Box struct {
	fullbox.Box
	FragmentDuration uint64
}

func New(i *box.Info) box.Boxed {
	return &Box{*fullbox.New(i), 0}
}

func (b Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (Box) Type() string {
	return MEHD
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)

	if b.Version == 0 {
		b.FragmentDuration = uint64(sr.Uint32())
	} else {
		b.FragmentDuration = sr.Uint64()
	}

	return box.FullRead(len(src))
}
