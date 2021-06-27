package alst

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/base"
	"github.com/jwhittle933/streamline/media/mp4/box"
)

const (
	ALST = "alst"
)

type Box struct {
	base.Box
	RollCount        uint16
	FistOutputSample uint16
	SampleOffset     []uint32
	NumOutputSamples []uint16
	NumTotalSample   []uint16
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		make([]uint32, 0, 0),
		make([]uint16, 0, 0),
		make([]uint16, 0, 0),
	}
}

func (Box) Type() string {
	return ALST
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s, rollcount=%d, firstoutputsample=%d, sampleoffset=%+v, numoutputsamples=%+v, numtotalsamples=%+v",
		b.Info(),
		b.RollCount,
		b.FistOutputSample,
		b.SampleOffset,
		b.NumOutputSamples,
		b.NumTotalSample,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.RollCount = sr.Uint16()
	b.FistOutputSample = sr.Uint16()
	b.SampleOffset = make([]uint32, int(b.RollCount))
	for i := 0; i < len(b.SampleOffset); i++ {
		b.SampleOffset[i] = sr.Uint32()
	}

	rem := sr.Remaining()
	if len(rem)/4 <= 0 {
		return box.FullRead(len(src))
	}

	b.NumOutputSamples = make([]uint16, len(rem))
	b.NumTotalSample = make([]uint16, len(rem))
	for i := 0; i < len(rem); i++ {
		b.NumOutputSamples[i] = sr.Uint16()
		b.NumTotalSample[i] = sr.Uint16()
	}

	return box.FullRead(len(src))
}
