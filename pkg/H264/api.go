package h264

import (
	"bytes"
	"fmt"
	"io"
)

const (
	IFrameByte byte = iota
	PFrameByte
	BFrameByte
)

const (
	NALU_TYPE_NOT_DEFINED byte = iota
	NALU_TYPE_SLICE
	NALU_TYPE_DPA
	NALU_TYPE_DPB
	NALU_TYPE_DPC
	NALU_TYPE_IDR
	NALU_TYPE_SEI
	NALU_TYPE_SPS
	NALU_TYPE_PPS
	NALU_TYPE_AUD
	NALU_TYPE_EOESQ
	NALU_TYPE_EOSTREAM
	NALU_TYPE_FILLER
)

const (
	NALU_BYTES_LEN  int = 4
	MAX_SPS_PPS_LEN int = 2 * 1024
)

var (
	startCode = []byte{0x00, 0x00, 0x00, 0x01}
	naluAud   = []byte{0x00, 0x00, 0x00, 0x01, 0x09, 0xf0}
)

type sequenceHeader struct {
	configVersion        byte
	avcProfileIndication byte
	profileCompatibility byte
	avcLevelIndication   byte
	reserved1            byte
	naluLen              byte
	reserved2            byte
	spsNum               byte
	ppsNum               byte
	spsLen               int
	ppsLen               int
}

type Parser struct {
	frameType    byte
	specificInfo []byte
	pps          *bytes.Buffer
}

func New() *Parser {
	return &Parser{pps: bytes.NewBuffer(make([]byte, MAX_SPS_PPS_LEN))}
}

func (p *Parser) parseSpecificInfo(src []byte) error {
	if len(src) < 9 {
		return fmt.Errorf("dec data is nil")
	}


	seq := sequenceHeader{
		configVersion:        src[0],
		avcProfileIndication: src[1],
		profileCompatibility: src[2],
		avcLevelIndication:   src[3],
		reserved1:            src[4] & 0xfc,
		naluLen:              src[4]&0x03 + 1,
		reserved2:            src[5] >> 5,
		spsNum:               src[5] & 0x1f,
		spsLen:               int(src[6]<<8) | int(src[7]),
	}

	if len(src[8:]) < seq.spsLen || seq.spsLen <= 0 {
		return fmt.Errorf("sps data error")
	}

	sps := append(startCode, src[8:(8+seq.spsLen)]...)

	tmpBuf := src[(8 + seq.spsLen):]
	if len(tmpBuf) < 4 {
		return fmt.Errorf("pps header error")
	}

	seq.ppsNum = tmpBuf[0]
	seq.ppsLen = 0<<16 | int(tmpBuf[1])<<8 | int(tmpBuf[2])

	if len(tmpBuf[3:]) < seq.ppsLen || seq.ppsLen <= 0 {
		return fmt.Errorf("pps data error")
	}

	pps := append(startCode, tmpBuf[3:]...)

	p.specificInfo = append(
		p.specificInfo,
		append(sps, pps...)...,
	)

	return nil
}

func (p *Parser) isNaluHeader(src []byte) bool {
	if len(src) < NALU_BYTES_LEN {
		return false
	}

	return src[0] == 0x00 &&
		src[1] == 0x00 &&
		src[2] == 0x00 &&
		src[3] == 0x01
}

func (p *Parser) naluSize(src []byte) (int, error) {
	if len(src) < NALU_BYTES_LEN {
		return 0, fmt.Errorf("nalusizedata is invalid")
	}

	buf := src[:NALU_BYTES_LEN]
	var size int

	for i := 0; i < len(buf); i++ {
		size = size<<8 + int(buf[i])
	}

	return size, nil
}

func (p *Parser) getAnnexbH264(src []byte, w io.Writer) error {
	dataSize := len(src)
	if dataSize < NALU_BYTES_LEN {
		return fmt.Errorf("video data did not match")
	}

	p.pps.Reset()

	_, err := w.Write(naluAud)
	if err != nil {
		return err
	}

	index := 0
	nalLen := 0
	hasSpsPps := false
	hasWriteSpsPps := false

	for dataSize > 0 {
		nalLen, err = p.naluSize(src)
		if err != nil {
			return err
		}

		index += NALU_BYTES_LEN
		dataSize -= NALU_BYTES_LEN

		if dataSize >= nalLen && len(src[index:]) >= nalLen && nalLen > 0 {
			nalType := src[index] & 0x1f
			switch nalType {
			case NALU_TYPE_AUD:
				fallthrough
			case NALU_TYPE_IDR:
				if !hasWriteSpsPps {
					hasWriteSpsPps = true
					if !hasSpsPps {
						if _, err := w.Write(p.specificInfo); err != nil {
							return err
						}
					} else {
						if _, err := w.Write(p.pps.Bytes()); err != nil {
							return err
						}
					}
				}

				fallthrough
			case NALU_TYPE_SLICE:
				fallthrough
			case NALU_TYPE_SEI:
				if _, err := w.Write(append(startCode, src[index:index+nalLen]...)); err != nil {
					return err
				}
			case NALU_TYPE_SPS:
				fallthrough
			case NALU_TYPE_PPS:
				hasSpsPps = true
				if _, err := p.pps.Write(append(startCode, src[index:index+nalLen]...)); err != nil {
					return err
				}
			}

			index += nalLen
			dataSize -= nalLen
		} else {
			return fmt.Errorf("body length error")
		}
	}

	return nil
}

func (p *Parser) Parse(b []byte, isSeq bool, w io.Writer) error {
	var err error

	switch isSeq {
	case true:
		err = p.parseSpecificInfo(b)
	case false:
		if p.isNaluHeader(b) {
			_, err = w.Write(b)
		} else {
			err = p.getAnnexbH264(b, w)
		}
	}


	return err
}
