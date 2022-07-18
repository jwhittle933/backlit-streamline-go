// Package tfdt (Track Fragment Base Media Decoder Time)
package tfdt

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/fullbox"
)

const (
	TFDT string = "tfdt"
)

type Box struct {
	fullbox.Box
	BaseMediaDecodeTime uint64
}

func New(i *box.Info) box.Boxed {
	return &Box{
		*fullbox.New(i),
		0,
	}
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s, basemediadecodetime=%d",
		b.Info(),
		b.BaseMediaDecodeTime,
	)
}

func (Box) Type() string {
	return TFDT
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)
	if b.Version == 0 {
		b.BaseMediaDecodeTime = uint64(sr.Uint32())
	} else {
		b.BaseMediaDecodeTime = sr.Uint64()
	}

	return box.FullRead(len(src))
}
