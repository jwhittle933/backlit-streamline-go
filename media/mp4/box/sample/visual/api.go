package visual

import (
	"bytes"
	"fmt"
	slicereader2 "github.com/jwhittle933/streamline/bits/slicereader"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	children2 "github.com/jwhittle933/streamline/media/mp4/box/children"
	avcC2 "github.com/jwhittle933/streamline/media/mp4/box/sample/visual/avcC"
	btrt2 "github.com/jwhittle933/streamline/media/mp4/box/sample/visual/btrt"
	clap2 "github.com/jwhittle933/streamline/media/mp4/box/sample/visual/clap"
	hvcC2 "github.com/jwhittle933/streamline/media/mp4/box/sample/visual/hvcC"
	pasp2 "github.com/jwhittle933/streamline/media/mp4/box/sample/visual/pasp"
	scanner2 "github.com/jwhittle933/streamline/media/mp4/box/scanner"
)

var (
	Children = children2.Registry{
		avcC2.AVCC: avcC2.New,
		hvcC2.HVCC: hvcC2.New,
		btrt2.BTRT: btrt2.New,
		clap2.CLAP: clap2.New,
		pasp2.PASP: pasp2.New,
	}
)

type SampleEntry struct {
	base2.Box
	DataReferenceIndex uint16
	Width              uint16
	Height             uint16
	HorizResolution    uint32
	VertResolution     uint32
	FrameCount         uint16
	CompressorName     string
	Children           []box2.Boxed
}

func New(i *box2.Info) box2.Boxed {
	return &SampleEntry{
		base2.Box{BoxInfo: i},
		0,
		0,
		0,
		0,
		0,
		0,
		"",
		make([]box2.Boxed, 0, 0),
	}
}

func (s SampleEntry) String() string {
	out := fmt.Sprintf(
		"%s, data_reference_index=%d, width=%d, height=%d, horiz_res=%d, vert_res=%d, frames=%d, compressorname=%s children=%d,",
		s.Info(),
		s.DataReferenceIndex,
		s.Width,
		s.Height,
		s.HorizResolution,
		s.VertResolution,
		s.FrameCount,
		s.CompressorName,
		len(s.Children),
	)

	for _, child := range s.Children {
		out += fmt.Sprintf("\n------------->%s", child)
	}

	return out
}

func (s SampleEntry) Type() string {
	return s.Info().Type.String()
}

func (s *SampleEntry) Write(src []byte) (int, error) {
	sr := slicereader2.New(src)
	sr.Skip(6) // 6 bytes for reserved
	s.DataReferenceIndex = sr.Uint16()

	sr.Skip(4)  // 4 bytes predefined/reserved
	sr.Skip(12) // 3x32 bits for predefined
	s.Width = sr.Uint16()
	s.Height = sr.Uint16()
	s.HorizResolution = sr.Uint32()
	s.VertResolution = sr.Uint32()

	sr.Uint32() // skip 32 bits for reserved
	s.FrameCount = sr.Uint16()
	s.WriteCompressor(sr)

	sr.Skip(2) // 16 bits for depth == 0x0018
	sr.Skip(2) // 16 bits for reserved

	remaining := bytes.NewReader(sr.Remaining())
	sc := scanner2.New(remaining)

	var err error
	s.Children, err = sc.ScanAllChildren(Children)
	if err != nil {
		return 0, err
	}

	return box2.FullRead(len(src))
}

// WriteCompressor TODO: write compressor name
func (s *SampleEntry) WriteCompressor(sr *slicereader2.Reader) {
	compressorNameLength := sr.Uint8()
	if compressorNameLength > 31 {
		s.CompressorName = "INVALID"
	}

	s.CompressorName = sr.String(int(compressorNameLength))
	if s.CompressorName == "" {
		s.CompressorName = "EMPTY"
	}
	sr.Skip(int(31 - compressorNameLength))
	//sr.Skip(32)
}
