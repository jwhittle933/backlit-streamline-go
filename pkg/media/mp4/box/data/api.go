package data

import (
	"fmt"
	"strings"
	"unicode"
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
	EncodingType Encoding
	Language     uint32
	Data         []byte
}

type Encoding uint32

func (Box) Type() string {
	return DATA
}

func (b Box) String() string {
	if uint32(b.EncodingType) == StringUTF8 {
		return fmt.Sprintf("\"%s\"", strings.Map(func(r rune) rune {
			if unicode.IsGraphic(r)	 {
				return r
			}

			return rune('.')
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
