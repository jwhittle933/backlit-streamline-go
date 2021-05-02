package moov

import (
	"fmt"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/children"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/mvhd"
)

const (
	MOOV string = "moov"
)

var moovChildren = children.Registry{
	mvhd.MVHD: mvhd.New,
	//trak.TRAK: trak.New,
	//tkhd.TKHD: tkhd.New,
	//edts.EDTS: edts.New,
	//elst.ELST: elst.New,
	//mdia.MDIA: mdia.New,
	//mdhd.MDHD: mdhd.New,
	//hdlr.HDLR: hdlr.New,
	//minf.MINF: minf.New,
	//vmhd.VMHD: vmhd.New,
	//dinf.DINF: dinf.New,
	//dref.DREF: dref.New,
	//url.URL:   url.New,
	//stbl.STBL: stbl.New,
	//stsd.STSD: stsd.New,
	//avc1.AVC1: avc1.New,
	//avcC.AVCC: avcC.New,
	//pasp.PASP: pasp.New,
	//stts.STTS: stts.New,
	//ctts.CTTS: ctts.New,
	//stsc.STSC: stsc.New,
	//stsz.STSZ: stsz.New,
	//stco.STCO: stco.New,
}

type Box struct {
	base.Box
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}}
}

func (Box) Type() string {
	return MOOV
}

func (b Box) String() string {
	return fmt.Sprintf(
		"[%s] hex=%s, offset=%d, size=%d, header=%d",
		b.BoxInfo.Type.String(),
		b.BoxInfo.Type.HexString(),
		b.BoxInfo.Offset,
		b.BoxInfo.Size,
		b.BoxInfo.HeaderSize,
	)
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	// iteratively parse children
	return len(src), nil
}
