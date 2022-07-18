package cslg

import (
	"fmt"
	slicereader2 "github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/base"
	box2 "github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	CSLG string = "cslg"
)

type Box struct {
	base.Box
	Version                      byte
	Flags                        uint32
	CompositionToDTSShift        int64
	LeastDecodeToDisplayDelta    int64
	GreatestDecodeToDisplayDelta int64
	CompositionStartTime         int64
	CompositionEndTime           int64
}

func New(i *box2.Info) box2.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		0,
		0,
		0,
		0,
	}
}

func (b *Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (Box) Type() string {
	return CSLG
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader2.New(src)
	versionFlags := sr.Uint32()
	b.Version = byte(versionFlags >> 24)
	b.Flags = versionFlags & 0x00ffffff

	if b.Version == 0 {
		b.CompositionToDTSShift = int64(sr.Uint32())
		b.LeastDecodeToDisplayDelta = int64(sr.Uint32())
		b.GreatestDecodeToDisplayDelta = int64(sr.Uint32())
		b.CompositionStartTime = int64(sr.Uint32())
		b.CompositionEndTime = int64(sr.Uint32())
	} else {
		b.CompositionToDTSShift = sr.Int64()
		b.LeastDecodeToDisplayDelta = sr.Int64()
		b.GreatestDecodeToDisplayDelta = sr.Int64()
		b.CompositionStartTime = sr.Int64()
		b.CompositionEndTime = sr.Int64()
	}

	return box2.FullRead(len(src))
}
