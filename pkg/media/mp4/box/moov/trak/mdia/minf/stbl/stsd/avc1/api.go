package avc1

import (
	"encoding/binary"
	"fmt"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/sample"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/children"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsd/avcC"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsd/pasp"
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
	sample.Audio
	AuxInfoType           [4]byte
	AuxInfoTypeParameter  uint32
	DefaultSampleInfoSize uint8
	SampleCount           uint32
	SampleInfoSize        []uint8
}

func New(i *box.Info) box.Boxed {
	return &Box{
		sample.Audio{
			Entry: sample.Entry{
				Box: base.Box{BoxInfo: i},
				DataReferenceIndex: 0,
			},
			Version: 0,
			ChannelCount: 0,
			SampleSize: 0,
			Predefined: 0,
			SampleRate: 0,
			QuickTimeData: make([]byte, 0),
		},
		[4]byte{},
		0,
		0,
		0,
		make([]uint8, 0),
	}
}

func (Box) Type() string {
	return AVC1
}

func (b *Box) String() string {
	return b.Info().String() + fmt.Sprintf(
		", aux_info_type=%s, aux_parameter=%d",
		b.AuxInfoType,
		b.AuxInfoTypeParameter,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	//fmt.Println("avc1:", string(src[:]))
	copy(b.AuxInfoType[:], src[0:4])
	b.AuxInfoTypeParameter = binary.BigEndian.Uint32(src[4:8])
	b.DefaultSampleInfoSize = src[8]

	return len(src), nil
}
