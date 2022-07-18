package data

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	DATA               string = "data"
	Binary                    = 0
	StringUTF8                = 1
	StringUTF16               = 2
	StringMac                 = 3
	StringJPEG                = 14
	SignedIntBigEndian        = 21
	Float32BigEndian          = 22
	Float64BigEndian          = 23
)

type Box struct {
	base.Box
	EncodingType Encoding
	Language     uint32
	Data         []byte
}

type Encoding uint32

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		make([]byte, 0),
	}
}

func (Box) Type() string {
	return DATA
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s, encoding=%s, data=%s",
		b.Info(),
		b.EncodingType,
		b.DataString(),
	)
}

func (b Box) DataString() string {
	if uint32(b.EncodingType) == StringUTF8 {
		return fmt.Sprintf("\"%s\"", strings.Map(func(r rune) rune {
			if unicode.IsGraphic(r) {
				return r
			}

			return '.'
		}, string(b.Data)))
	}

	return "data not UTF8-encoded. Cannot print"
}

func (e Encoding) String() string {
	switch uint32(e) {
	case Binary:
		return "BINARY"
	case StringUTF8:
		return "UTF8"
	case StringUTF16:
		return "UTF16"
	case StringMac:
		return "MAC_STR"
	case StringJPEG:
		return "JPEG"
	case SignedIntBigEndian:
		return "INT"
	case Float32BigEndian:
		return "FLOAT32"
	case Float64BigEndian:
		return "FLOAT64"
	}

	return "Unknown"
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)

	b.EncodingType = Encoding(sr.Uint32())
	b.Language = sr.Uint32()
	b.Data = sr.Remaining()

	return box.FullRead(len(src))
}
