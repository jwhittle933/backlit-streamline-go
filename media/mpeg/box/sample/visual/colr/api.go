package colr

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits"
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	COLR string = "colr"
)

type Box struct {
	base.Box
	ColorType               string
	ColorPrimaries          uint16
	TransferCharacteristics uint16
	MatrixCoefficients      uint16
	FullRangeFlag           bool
	RestrictedProfile       []byte
	UnrestrictedProfile     []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		"",
		0,
		0,
		0,
		false,
		nil,
		nil,
	}
}

func (Box) Type() string {
	return COLR
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s, colortype=%s, colorprimaries=%d, transfercharacteristics=%d, matrix_coefficients=%d, fullrangeflag=%+v, restricted=%+v, unrestricted=%+v",
		b.Info(),
		b.ColorType,
		b.ColorPrimaries,
		b.TransferCharacteristics,
		b.MatrixCoefficients,
		b.FullRangeFlag,
		b.RestrictedProfile,
		b.UnrestrictedProfile,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.ColorType = sr.String(4)

	if b.ColorType == "nclx" {
		b.ColorPrimaries = sr.Uint16()
		b.TransferCharacteristics = sr.Uint16()
		b.MatrixCoefficients = sr.Uint16()

		br := bits.NewReader(sr.Slice(1))
		bit, _ := br.ReadBit()
		b.FullRangeFlag = bit.Bool()
		// 7 bits reserved
	} else if b.ColorType == "rICC" {
		b.RestrictedProfile = sr.Remaining()
	} else if b.ColorType == "prof" {
		b.UnrestrictedProfile = sr.Remaining()
	}

	return box.FullRead(len(src))
}
