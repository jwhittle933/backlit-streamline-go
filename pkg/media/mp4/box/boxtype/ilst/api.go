package ilst

const (
	ILST     string = "ilst"
	IlstDash string = "----"
	IlstaART string = "aART"
	IlstAkID string = "akID"
	IlstApID string = "apID"
	IlstAtID string = "atID"
	IlstCmID string = "cmID"
	IlstCovr string = "covr"
	IlstCpil string = "cpil"
	IlstCprt string = "cprt"
	IlstDesc string = "desc"
	IlstDisk string = "disk"
	IlstEgid string = "egid"
	IlstGeID string = "geID"
	IlstGnre string = "gnre"
	IlstPcst string = "pcst"
	IlstPgap string = "pgap"
	IlstPlID string = "plID"
	IlstPurd string = "purd"
	IlstPurl string = "purl"
	IlstRtng string = "rtng"
	IlstSfID string = "sfid"
	IlstSoaa string = "soaa"
	IlstSoal string = "soal"
	IlstSoar string = "soar"
	IlstSoco string = "soco"
	IlstSonm string = "sonm"
	IlstSosn string = "sosn"
	IlstStik string = "stik"
	IlstTmpo string = "tmpo"
	IlstTrkn string = "trkn"
	IlstTven string = "tven"
	IlstTves string = "tves"
	IlstTvnn string = "tvnn"
	IlstTvsh string = "tvsh"
	IlstTvsn string = "tvsn"
)

var (
	IlstART  string = string([]byte{0xA9, 'A', 'R', 'T'})
	IlstAlb  string = string([]byte{0xA9, 'a', 'l', 'b'})
	IlstCmt  string = string([]byte{0xA9, 'c', 'm', 't'})
	IlstCom  string = string([]byte{0xA9, 'c', 'o', 'm'})
	IlstDay  string = string([]byte{0xA9, 'd', 'a', 'y'})
	IlstGen  string = string([]byte{0xA9, 'g', 'e', 'n'})
	IlstGrp  string = string([]byte{0xA9, 'g', 'r', 'p'})
	IlstNam  string = string([]byte{0xA9, 'n', 'a', 'm'})
	IlstToo  string = string([]byte{0xA9, 't', 'o', 'o'})
	IlstWrt  string = string([]byte{0xA9, 'w', 'r', 't'})
)

type Box struct {
	//
}

func (Box) Type() string {
	return ILST
}
