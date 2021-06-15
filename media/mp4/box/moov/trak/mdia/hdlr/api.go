// Package hdlr (Handler) declares the media handler type
package hdlr

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/fullbox"
)

const (
	HDLR string = "hdlr"
)

// Box is trak/mdia/hdlr box type
type Box struct {
	fullbox.Box
	// PreDefined: component_type of QuickTime
	// pre_defined of ISO-14496 always has 0
	// component_type has mhlr or dhlr
	HandlerType   [4]byte
	Name          string
	LacksNullTerm bool
}

func New(i *box.Info) box.Boxed {
	return &Box{
		*fullbox.New(i),
		[4]byte{},
		"",
		false,
	}
}

func (Box) Type() string {
	return HDLR
}

func (b *Box) String() string {
	return fmt.Sprintf(
		"%s, version=%d, flags=%d, handlertype=%s, name=%s, lacksnullterm=%+v",
		b.Info(),
		b.Version,
		b.Flags,
		b.HandlerType,
		b.Name,
		b.LacksNullTerm,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)
	sr.Skip(4)

	copy(b.HandlerType[:], sr.Slice(4))
	// skip 12 reserved
	sr.Skip(12)
	if len(src) > 24 {
		final := len(src) - 1
		last := src[final]
		if last != 0 {
			final++
			b.LacksNullTerm = true
		}

		b.Name = string(src[24:final])
	} else {
		b.LacksNullTerm = true
	}

	return box.FullRead(len(src))
}
