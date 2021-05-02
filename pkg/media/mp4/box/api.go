// Package box defines functions to read from
// and write to Boxes within in ISO BMFF box/atom
package box

import (
	"fmt"
	"io"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box/boxtype"
)

const (
	SmallHeader = 8
	LargeHeader = 16
)

type Typed interface {
	Type() string
}

type Informed interface {
	Info() *Info
}

type Boxed interface {
	io.Writer
	fmt.Stringer
	Typed
	Informed
}

type Info struct {
	Offset      uint64
	Size        uint64
	Type        boxtype.BoxType
	HeaderSize  uint64
	ExtendToEOF bool
}

func SeekPayload(s io.Seeker, b Boxed) (int64, error) {
	info := b.Info()
	return s.Seek(int64(info.Offset+info.HeaderSize), io.SeekStart)
}
