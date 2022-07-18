package hvc

import (
	"errors"
	"fmt"
	slicereader2 "github.com/jwhittle933/streamline/bits/slicereader"
	box2 "github.com/jwhittle933/streamline/media/mpeg/box"
)

var ErrorLengthSize = errors.New("can only handle 4byte NALU length size")

type DecoderConfig struct {
	ConfigurationVersion             byte
	GeneralProfileSpace              byte
	GeneralTierFlag                  bool
	GeneralProfileIDC                byte
	GeneralProfileCompatibilityFlags uint32
	GeneralConstraintIndicatorFlags  uint64
	GeneralLevelIDC                  byte
	MinSpatialSegmentationIDC        uint16
	ParallellismType                 byte
	ChromaFormatIDC                  byte
	BitDepthLumaMinus8               byte
	BitDepthChromaMinus8             byte
	AvgFrameRate                     uint16
	ConstantFrameRate                byte
	NumTemporalLayers                byte
	TemporalIDNested                 byte
	LengthSizeMinusOne               byte
	NaluArrays                       []NaluArray
}

type NaluArray struct {
	completeAndType byte
	NALUs           [][]byte
}

func New() DecoderConfig {
	return DecoderConfig{}
}

func (d *DecoderConfig) Write(src []byte) (int, error) {
	sr := slicereader2.New(src)
	d.ConfigurationVersion = sr.Uint8()
	if d.ConfigurationVersion != 1 {
		return 0, fmt.Errorf("HEVC decoder configuration record version %d unknown", d.ConfigurationVersion)
	}

	next8 := sr.Uint8()
	d.GeneralProfileSpace = (next8 >> 6) & 0x3
	d.GeneralTierFlag = (next8>>5)&0x01 == 0x1
	d.GeneralLevelIDC = next8 & 0x1f

	d.GeneralProfileCompatibilityFlags = sr.Uint32()
	d.GeneralConstraintIndicatorFlags = (uint64(sr.Uint32()) << 16) | uint64(sr.Uint16())
	d.GeneralLevelIDC = sr.Uint8()
	d.MinSpatialSegmentationIDC = sr.Uint16() & 0x0fff
	d.ParallellismType = sr.Uint8() & 0x3
	d.ChromaFormatIDC = sr.Uint8() & 0x3
	d.BitDepthLumaMinus8 = sr.Uint8() & 0x7
	d.BitDepthChromaMinus8 = sr.Uint8() & 0x7
	d.AvgFrameRate = sr.Uint16()

	next8 = sr.Uint8()
	d.ConstantFrameRate = (next8 >> 6) & 0x3
	d.NumTemporalLayers = (next8 >> 3) & 0x7
	d.TemporalIDNested = (next8 >> 2) & 0x1
	d.LengthSizeMinusOne = next8 & 0x3
	if d.LengthSizeMinusOne != 3 {
		return 0, ErrorLengthSize
	}

	numArrays := int(sr.Uint8())
	d.NaluArrays = make([]NaluArray, 0, numArrays)
	for i := 0; i < len(d.NaluArrays); i++ {

		arr := NaluArray{completeAndType: sr.Uint8()}
		numNALUs := int(sr.Uint16())
		arr.NALUs = make([][]byte, 0, numNALUs)
		for j := 0; j < len(arr.NALUs); j++ {
			length := int(sr.Uint16())
			arr.NALUs = append(arr.NALUs, sr.Slice(length))
		}

		d.NaluArrays[i] = arr
	}

	return box2.FullRead(len(src))
}
