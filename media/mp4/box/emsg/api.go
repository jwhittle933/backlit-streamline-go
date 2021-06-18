package emsg

import (
	"fmt"
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/fullbox"
)

const (
	EMSG string = "emsg"
)

type Box struct {
	fullbox.Box
	Timescale             uint32
	PresentationTimeDelta uint32
	PresentationTime      uint64
	EventDuration         uint32
	Id                    uint32
	SchemeIdUri           string
	Value                 string
}

func New(i *box.Info) box.Boxed {
	return &Box{
		*fullbox.New(i),
		0,
		0,
		0,
		0,
		0,
		"",
		"",
	}
}

func (Box) Type() string {
	return EMSG
}

func (b Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)

	if b.Version == 1 {
		b.Timescale = sr.Uint32()
		b.PresentationTime = sr.Uint64()
		b.EventDuration = sr.Uint32()
		b.Id = sr.Uint32()
		b.SchemeIdUri = sr.NullTermString()
		b.Value = sr.NullTermString()
	} else if b.Version == 0 {
		b.SchemeIdUri = sr.NullTermString()
		b.Value = sr.NullTermString()
		b.Timescale = sr.Uint32()
		b.PresentationTimeDelta = sr.Uint32()
		b.EventDuration = sr.Uint32()
		b.Id = sr.Uint32()
	} else  {
		return 0, fmt.Errorf("unknown version %d for emsg", b.Version)
	}


	return box.FullRead(len(src))
}
