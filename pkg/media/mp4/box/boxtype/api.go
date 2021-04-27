// Package boxtype
// see https://github.com/abema/go-mp4/blob/8afbbcf4788ca39d4147047bce9dc915b013c9fa/box_types.go#L253
package boxtype

import "fmt"

type BoxType [4]byte

const (
	EMSG string = "emsg"
	ESDS string = "esds"
	FREE string = "free"
	SKIP string = "skip"
	FRMA string = "frma"
	FTYP string = "ftyp"
	HDLR string = "hdlr"
	ILSD string = "ilsd"
	DATA string = "data"
	MDAT string = "mdat"
	MDHD string = "mdhd"
	MDIA string = "mdia"
	MEHD string = "mehd"
	META string = "meta"
	MFHD string = "mfhd"
	MFRA string = "mfra"
	MFRO string = "mfro"
	MINF string = "minf"
	MOOF string = "moof"
	MOOV string = "moov"
	MVEX string = "mvex"
	MVHD string = "mvhd"
	PSSH string = "pssh"
	SAIO string = "saio"
	SAIZ string = "saiz"
	AVC1 string = "avc1"
	ENCV string = "encv"
	MP4A string = "mp4a"
	ENCA string = "enca"
	AVCC string = "avcC"
	PASP string = "pasp"
	SBGP string = "sbgp"
	SCHI string = "schi"
	SCHM string = "schm"
	SDTP string = "sdtp"
	SGPD string = "sgpd"
	SIDX string = "sidx"
	SINF string = "sing"
	SMHD string = "smhd"
	STBL string = "stbl"
	STCO string = "stco"
	STSC string = "stsc"
	STSD string = "stsd"
	STSS string = "stss"
	STSZ string = "stsz"
	STTS string = "stts"
	STYP string = "styp"
	TENC string = "tenc"
	TFDT string = "tfdt"
	TFHD string = "tfhd"
	TFRA string = "tfra"
	TKHD string = "tkhd"
	TRAF string = "traf"
	TRAK string = "trak"
	TREP string = "trep"
	TREX string = "trex"
	TRUN string = "trun"
	VMHD string = "vmhd"
	WAVE string = "wave"
)

func New(code string) BoxType {
	if len(code) != 4 {
		panic(fmt.Errorf("invalid box type id length [%s]", code))
	}

	return BoxType{code[0], code[1], code[2], code[3]}
}

func (b BoxType) String() string {
	return fmt.Sprintf("0x%02x%02x%02x%02x", b[0], b[1], b[2], b[3])
}
func isASCII(c byte) bool {
	return c >= 0x20 && c <= 0x7e
}

func isPrintable(c byte) bool {
	return isASCII(c) || c == 0xa9
}
