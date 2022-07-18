package avcC

import (
	"fmt"

	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/box/sample/visual/avc"
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
	base.Box
	avc.DecoderConfig
}

type ParameterSet struct {
	Length  uint16
	NALUnit []byte
}

type NALUnit struct {
	_forbidden uint8
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, avc.New()}
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
