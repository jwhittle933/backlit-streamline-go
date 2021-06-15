// Package stts (Time to Sample) decoding
package stts

import (
	"encoding/binary"
	"fmt"
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/fullbox"
)

const (
	STTS string = "stts"
)

type Box struct {
	fullbox.Box
	SampleCount     []uint32
	SampleTimeDelta []uint32
}

type Entry struct {
	SampleCount uint32
	SampleDelta uint32
}

func New(i *box.Info) box.Boxed {
	return &Box{*fullbox.New(i), make([]uint32, 0, 0), make([]uint32, 0, 0)}
}

func (Box) Type() string {
	return STTS
}

func (b *Box) String() string {
	return fmt.Sprintf(
		"%s, samplecount=%d, samplecounts=%+v, sampledeltas=%+v",
		b.Info(),
		len(b.SampleCount),
		b.SampleCount,
		b.SampleTimeDelta,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)

	count := sr.Uint32()
	b.SampleCount = make([]uint32, 0, count)
	b.SampleTimeDelta = make([]uint32, 0, count)
	for i := 0; i < len(b.SampleCount); i++ {
		b.SampleCount[i] = binary.BigEndian.Uint32(src[(8 + 8*i):(12 + 8*i)])
		b.SampleTimeDelta[i] = binary.BigEndian.Uint32(src[(12 + 8*i):(16 + 8*i)])
	}

	return box.FullRead(len(src))
}
