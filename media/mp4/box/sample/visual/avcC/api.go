package avcC

import (
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	avc2 "github.com/jwhittle933/streamline/media/mp4/box/sample/visual/avc"
)

const (
	AVCC            string = "avcC"
	BaselineProfile uint8  = 66
	MainProfile     uint8  = 77
	ExtendedProfile uint8  = 88
	HighProfile     uint8  = 100
	High10Profile   uint8  = 110
	High422Profile  uint8  = 122
)

type Box struct {
	base2.Box
	avc2.DecoderConfig
}

type ParameterSet struct {
	Length  uint16
	NALUnit []byte
}

type NALUnit struct {
	_forbidden uint8
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, avc2.New()}
}

func (Box) Type() string {
	return AVCC
}

func (b *Box) String() string {
	return b.Info().String() + fmt.Sprintf(
		", profile=%d, compatibility=%d, level=%d, length_sps=%d, len_pps=%d, chroma_format=%d, bit_depth_luma=%d, bit_depth_chroma=%d, trailing_info=%+v",
		b.AVCProfileIndication,
		b.ProfileCompatibility,
		b.AVCLevelIndication,
		len(b.SpsNALUs),
		len(b.PpsNALUs),
		b.ChromaFormat,
		b.BitDepthLumaMinus1,
		b.BitDepthChromaMinus1,
		b.NoTrailingInfo,
	)
}
