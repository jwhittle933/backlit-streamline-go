package loci

import (
	"fmt"
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/fullbox"
)

const (
	LOCI = "loci"
)

type Box struct {
	fullbox.Box
	Pad             byte
	Language        string
	Name            string
	Role            byte
	Long            int32
	Lat             int32
	Alt             int32
	AstrBody        string
	AdditionalNotes string
}

func New(i *box.Info) box.Boxed {
	return &Box{
		*fullbox.New(i),
		0,
		"",
		"",
		0,
		0,
		0,
		0,
		"",
		"",
	}
}

func (Box) Type() string {
	return LOCI
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s, version=%d, flags=%d, pad=%d, name=%s, role=%d, lon=%d, lat=%d, alt=%d, astrbody=%s, notes=%s",
		b.Info(),
		b.Version,
		b.Flags,
		b.Pad,
		//b.Language,
		b.Name,
		b.Role,
		b.Long,
		b.Lat,
		b.Alt,
		b.AstrBody,
		b.AdditionalNotes,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)

	padAndLang := sr.Uint16()
	b.Language = string([]byte{
		uint8(padAndLang&0x7c00>>10) + 0x60,
		uint8(padAndLang&0x03E0>>5) + 0x60,
		uint8(padAndLang&0x001F) + 0x06,
	})

	b.Name = sr.NullTermString()
	b.Role = sr.Uint8()
	b.Long = int32(sr.Uint32())
	b.Lat = int32(sr.Uint32())
	b.Alt = int32(sr.Uint32())
	b.AstrBody = sr.NullTermString()
	b.AdditionalNotes = sr.NullTermString()

	return box.FullRead(len(src))
}

func (b *Box) WriteLang(src []byte) {
	//
}
