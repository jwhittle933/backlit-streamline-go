package audio

import (
	"bytes"
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/base"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/sample/audio/esds"
	"github.com/jwhittle933/streamline/media/mp4/box/scanner"
	"github.com/jwhittle933/streamline/media/mp4/children"
)

var (
	Children = children.Registry{esds.ESDS: esds.New}
)

type SampleEntry struct {
	base.Box
	DataReferenceIndex uint16
	ChannelCount       uint16
	SampleSize         uint16
	SampleRate         uint16
	Children           []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &SampleEntry{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		0,
		make([]box.Boxed, 0),
	}
}

func (s SampleEntry) String() string {
	out := fmt.Sprintf(
		"%s, datareferenceindex=%d, channelcount=%d, samplesize=%d, samplerate=%d, children=%d",
		s.Info(),
		s.DataReferenceIndex,
		s.ChannelCount,
		s.SampleSize,
		s.SampleRate,
		len(s.Children),
	)

	for _, c := range s.Children {
		out += fmt.Sprintf("\n              %s", c)
	}

	return out
}

func (s SampleEntry) Type() string {
	return s.Info().Type.String()
}

func (s *SampleEntry) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	sr.Skip(6)

	s.DataReferenceIndex = sr.Uint16()
	sr.Skip(8)

	s.ChannelCount = sr.Uint16()
	s.SampleSize = sr.Uint16()
	sr.Skip(4)
	s.SampleRate = uint16(sr.Uint32() >> 16)

	sc := scanner.New(bytes.NewReader(sr.Remaining()))
	var err error
	s.Children, err = sc.ScanAllChildren(Children)
	if err != nil {
		return 0, err
	}

	return box.FullRead(len(src))
}
