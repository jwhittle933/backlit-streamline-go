package visual

import (
	"bytes"
	"fmt"
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
	"github.com/jwhittle933/streamline/media/mp4/box/children"
	"github.com/jwhittle933/streamline/media/mp4/box/sample/visual/avcC"
	"github.com/jwhittle933/streamline/media/mp4/box/sample/visual/btrt"
	"github.com/jwhittle933/streamline/media/mp4/box/sample/visual/clap"
	"github.com/jwhittle933/streamline/media/mp4/box/sample/visual/hvcC"
	"github.com/jwhittle933/streamline/media/mp4/box/sample/visual/pasp"
	"github.com/jwhittle933/streamline/media/mp4/box/scanner"
)

var (
	Children = children.Registry{
		avcC.AVCC: avcC.New,
		hvcC.HVCC: hvcC.New,
		btrt.BTRT: btrt.New,
		clap.CLAP: clap.New,
		pasp.PASP: pasp.New,
	}
)

type SampleEntry struct {
	base.Box
	DataReferenceIndex uint16
	Width              uint16
	Height             uint16
	HorizResolution    uint32
	VertResolution     uint32
	FrameCount         uint16
	CompressorName     string
	Children           []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &SampleEntry{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		0,
		0,
		0,
		"",
		make([]box.Boxed, 0, 0),
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
		out += fmt.Sprintf("\n              %s", child)
	}

	return out
}

func (s SampleEntry) Type() string {
	return s.Info().Type.String()
}

func (s *SampleEntry) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
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
	sc := scanner.New(remaining)

	var err error
	s.Children, err = sc.ScanAllChildren(Children)
	if err != nil {
		return 0, err
	}

	return box.FullRead(len(src))
}

// WriteCompressor TODO: write compressor name
func (s *SampleEntry) WriteCompressor(sr *slicereader.Reader) {
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
