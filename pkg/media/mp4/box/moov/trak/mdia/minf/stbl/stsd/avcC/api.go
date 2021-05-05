package avcC

import (
	"fmt"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	AVCC string = "avcC"
)

type Box struct {
	base.Box
	Version                    uint8
	Profile                    uint8
	ProfileCompatibility       uint8
	Level                      uint8
	_reserved                  uint8
	LengthSizeMinusOne         uint8
	_reserved2                 uint8
	SequenceParameterSetsLen   uint8
	SequenceParameterSets      []ParameterSet
	PictureParameterSetsLen    uint8
	PictureParameterSets       []ParameterSet
	HighProfileFieldsEnabled   bool
	_reserved3                 uint8
	ChromaFormat               uint8
	_reserved4                 uint8
	BitDepthLumaMinus8         uint8
	_reserved5                 uint8
	BitDepthChromaMinus8       uint8
	SequenceParameterSetExtLen uint8
	SequenceParameterSetExts   []ParameterSet
}

type ParameterSet struct {
	Length  uint16
	NALUnit []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{Box: base.Box{BoxInfo: i}}
}

func (Box) Type() string {
	return AVCC
}

func (b *Box) String() string {
	return b.Info().String() + fmt.Sprintf(
		", version=%d, profile=%d, profile_compat=%d, level=%d, length_minus_1=%d, seq_param_sets_len=%d",
		b.Version,
		b.Profile,
		b.ProfileCompatibility,
		b.Level,
		b.LengthSizeMinusOne,
		b.SequenceParameterSetsLen,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	b.Version = src[0]
	b.Profile = src[1]
	b.ProfileCompatibility = src[2]
	b.Level = src[3]
	// 4 reserved
	b.LengthSizeMinusOne = src[5]
	b.SequenceParameterSetsLen = src[7]
	b.SequenceParameterSets = make([]ParameterSet, b.SequenceParameterSetsLen)

	//offset := 7
	//for i := 0; uint8(i) < b.SequenceParameterSetsLen; i++ {
	//	ps := ParameterSet{
	//		Length: binary.BigEndian.Uint16(src[offset:offset+2]),
	//	}
	//	b.SequenceParameterSets[i] = ps
	//}

	return len(src), nil
}
