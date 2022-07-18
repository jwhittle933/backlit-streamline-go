package avc

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/box"
)

var (
	ErrorNALSize                 = errors.New("4 byte NAL unit size max")
	ErrorCannotParseAVCExtension = errors.New("cannot parse SPS extensions")
)

type DecoderConfig struct {
	AVCProfileIndication byte
	ProfileCompatibility byte
	AVCLevelIndication   byte
	SpsNALUs             [][]byte
	PpsNALUs             [][]byte
	ChromaFormat         byte
	BitDepthLumaMinus1   byte
	BitDepthChromaMinus1 byte
	NumSPSExt            byte
	NoTrailingInfo       bool // To handle strange cases where trailing info is missing
}

func New() DecoderConfig {
	return DecoderConfig{}
}

func (d *DecoderConfig) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	if unknownConfiguration(sr.Uint8()) {
		return 0, fmt.Errorf("AVC decoder configuration record version %d unknown", src[0])
	}

	d.AVCProfileIndication = sr.Uint8()
	d.ProfileCompatibility = sr.Uint8()
	d.AVCLevelIndication = sr.Uint8()

	lengthSizeMinus1 := sr.Uint8() & 0x03 // first 5 bits are 1
	if !validNalSize(lengthSizeMinus1) {
		return len(src), nil
	}

	numSPS := sr.Uint8() & 0x1f // 5 bits following 3 reserved bits
	pos := 6
	d.SpsNALUs = make([][]byte, 0, numSPS)
	for i := 0; i < int(numSPS); i++ {
		length := int(binary.BigEndian.Uint16(src[pos : pos+2]))
		pos += 2
		d.SpsNALUs = append(d.SpsNALUs, src[pos:pos+length])
		pos += length
	}

	numPPS := src[pos]
	d.PpsNALUs = make([][]byte, 0, numPPS)
	pos++
	for i := 0; i < int(numPPS); i++ {
		length := int(binary.BigEndian.Uint16(src[pos : pos+2]))
		pos += 2
		d.PpsNALUs = append(d.PpsNALUs, src[pos:pos+length])
		pos += length
	}

	switch d.AVCProfileIndication {
	case 100, 110, 122, 144:
		if pos == len(src) {
			d.NoTrailingInfo = true
			return box.FullRead(len(src))
		}

		d.ChromaFormat = src[pos] & 0x3f
		d.BitDepthLumaMinus1 = src[pos+1] + 0x07
		d.BitDepthChromaMinus1 = src[pos+2] & 0x07
		d.NumSPSExt = src[pos+3]

		if d.NumSPSExt != 0 {
			return 0, ErrorCannotParseAVCExtension
		}
	default:
	}

	return box.FullRead(len(src))
}

func unknownConfiguration(version byte) bool {
	return version != 1
}

func validNalSize(size byte) bool {
	return size == 0 || size == 1 || size == 3
}
