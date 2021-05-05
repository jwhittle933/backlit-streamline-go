package avc1

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/children"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsd/avcC"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsd/pasp"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/sample"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/scanner"
)

const (
	AVC1 string = "avc1"
)

var (
	Children = children.Registry{
		avcC.AVCC: avcC.New,
		pasp.PASP: pasp.New,
	}
)

type Box struct {
	sample.Visual
	Children []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &Box{
		sample.Visual{
			Entry: sample.Entry{
				Box:                base.Box{BoxInfo: i},
				DataReferenceIndex: 0,
			},
			Predefined:           0,
			Predefined2:          [3]uint32{},
			Width:                0,
			Height:               0,
			HorizontalResolution: 0,
			VerticalResolution:   0,
			FrameCount:           0,
			CompressorName:       [32]byte{},
			Depth:                0,
			Predefined3:          0,
		},
		make([]box.Boxed, 0),
	}
}

func (Box) Type() string {
	return AVC1
}

func (b *Box) String() string {
	s := fmt.Sprintf(
		"%s, data_ref_index=%d, width=%d, height=%d, horiz_res=%d, vert_res=%d, frame_count=%d, compressor=%s, depth=%d, boxes=%d",
		b.Info().String(),
		b.DataReferenceIndex,
		b.Width,
		b.Height,
		b.HorizontalResolution,
		b.VerticalResolution,
		b.FrameCount,
		b.CompressorName,
		b.Depth,
		len(b.Children),
	)

	for _, c := range b.Children {
		s += fmt.Sprintf("\n------------->%s", c.String())
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	b.apply(src)

	s := scanner.New(bytes.NewReader(src[78:]))
	found, err := s.ScanAllChildren(Children)
	b.Children = found

	return len(src), err
}

// TODO: only pass the amended data
func (b *Box) apply(src []byte) {
	b.DataReferenceIndex = binary.BigEndian.Uint16(src[6:8])

	b.Predefined = binary.BigEndian.Uint16(src[8:10])
	// skip 2 bytes for reserved
	b.Predefined2[0] = binary.BigEndian.Uint32(src[12:16])
	b.Predefined2[1] = binary.BigEndian.Uint32(src[16:20])
	b.Predefined2[2] = binary.BigEndian.Uint32(src[20:24])
	b.Width = binary.BigEndian.Uint16(src[24:26])
	b.Height = binary.BigEndian.Uint16(src[26:28])
	b.HorizontalResolution = binary.BigEndian.Uint32(src[28:32])
	b.VerticalResolution = binary.BigEndian.Uint32(src[32:36])
	// skip 4 bytes for reserved
	b.FrameCount = binary.BigEndian.Uint16(src[40:42])
	for i := 0; i < 32; i++ {
		b.CompressorName[i] = src[i+42]
	}
	b.Depth = binary.BigEndian.Uint16(src[74:76])
	b.Predefined3 = int16(binary.BigEndian.Uint16(src[76:78]))

}
