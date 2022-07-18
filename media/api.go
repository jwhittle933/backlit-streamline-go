// Package media parses and validates media formats
package media

import (
	"fmt"
	"io"
)

type Codec int

const (
	Unknown Codec = iota
	WAV
	ALAC
	FLAC
	APE
	OFR
	TAK
	WV
	TTA
	WMAL

	MP3
	M4A
	M4B
	M4P
	AAC
	OGG
	WMA
)

type Media interface {
	Valid() bool
	ReadAll() error
	Hex() string
	JSON() string
	Offset() (int64, error)
	io.ReadSeeker
	fmt.Stringer
}

func Open() {
	//
}
