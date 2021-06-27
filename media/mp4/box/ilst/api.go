package ilst

import (
	"bytes"
	"fmt"

	"github.com/jwhittle933/streamline/media/mp4/base"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/ilst/CToo"
	"github.com/jwhittle933/streamline/media/mp4/box/scanner"
	"github.com/jwhittle933/streamline/media/mp4/children"
)

const (
	ILST string = "ilst"
	Dash string = "----"
	ART  string = "aART"
	AkID string = "akID"
	ApID string = "apID"
	AtID string = "atID"
	CmID string = "cmID"
	Covr string = "covr"
	Cpil string = "cpil"
	Cprt string = "cprt"
	Desc string = "desc"
	Disk string = "disk"
	Egid string = "egid"
	GeID string = "geID"
	Gnre string = "gnre"
	Pcst string = "pcst"
	Pgap string = "pgap"
	PlID string = "plID"
	Purd string = "purd"
	Purl string = "purl"
	Rtng string = "rtng"
	SfID string = "sfid"
	Soaa string = "soaa"
	Soal string = "soal"
	Soar string = "soar"
	Soco string = "soco"
	Sonm string = "sonm"
	Sosn string = "sosn"
	Stik string = "stik"
	Tmpo string = "tmpo"
	Trkn string = "trkn"
	Tven string = "tven"
	Tves string = "tves"
	Tvnn string = "tvnn"
	Tvsh string = "tvsh"
	Tvsn string = "tvsn"
)

var (
	IlstART string = string([]byte{0xA9, 'A', 'R', 'T'})
	IlstAlb string = string([]byte{0xA9, 'a', 'l', 'b'})
	IlstCmt string = string([]byte{0xA9, 'c', 'm', 't'})
	IlstCom string = string([]byte{0xA9, 'c', 'o', 'm'})
	IlstDay string = string([]byte{0xA9, 'd', 'a', 'y'})
	IlstGen string = string([]byte{0xA9, 'g', 'e', 'n'})
	IlstGrp string = string([]byte{0xA9, 'g', 'r', 'p'})
	IlstNam string = string([]byte{0xA9, 'n', 'a', 'm'})
	IlstToo string = string([]byte{0xA9, 't', 'o', 'o'})
	IlstWrt string = string([]byte{0xA9, 'w', 'r', 't'})
)

var (
	Children = children.Registry{
		CToo.CTOO: CToo.New,
	}
)

type Box struct {
	base.Box
	Children []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, make([]box.Boxed, 0)}
}

func (Box) Type() string {
	return ILST
}

func (b Box) String() string {
	s := fmt.Sprintf("%s", b.Info())

	for _, c := range b.Children {
		s += fmt.Sprintf("\n        %s", c)
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	s := scanner.New(bytes.NewReader(src))

	var err error
	b.Children, err = s.ScanAllChildren(Children)
	if err != nil {
		return 0, err
	}

	return box.FullRead(len(src))
}
