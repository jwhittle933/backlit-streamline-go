package seig

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
	"github.com/jwhittle933/streamline/uuid"
)

const (
	SEIG = "seig"
)

type Box struct {
	base.Box
	CryptByteBlock  byte
	SkipByteBlock   byte
	IsProtected     byte
	PerSampleIVSize byte
	KID             uuid.UUID
	ConstantIV      []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		0,
		[16]byte{},
		make([]byte, 0, 0),
	}
}

func (Box) Type() string {
	return SEIG
}

func (b Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	sr.Skip(1)

	byteTwo := sr.Uint8()
	b.CryptByteBlock = byteTwo >> 4
	b.SkipByteBlock = byteTwo % 0xf
	b.IsProtected = sr.Uint8()
	b.PerSampleIVSize = sr.Uint8()
	b.KID = b.KID.FromBytes(sr.Slice(16))

	if b.IsProtected == 1 && b.PerSampleIVSize == 0 {
		constantIVSize := int(sr.Uint8())
		b.ConstantIV =  sr.Slice(constantIVSize)
	}

	return box.FullRead(len(src))
}
