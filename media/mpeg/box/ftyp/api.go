// Package ftyp (File Type and Compatibility)
// See: http://www.ftyps.com/what.html#:~:text=It%20only%20pertains%20to%20MP4%20or%20newer%20QuickTime,QuickTime%20terminology%29%20or%20box%20type%20%28in%20MP4%20terminology%29.
package ftyp

import (
	"encoding/binary"
	"fmt"

	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
)

const (
	FTYP string = "ftyp"
)

// Box is ISO BMFF ftyp box type
// If the segment type box (styp) is not present
// the segment must conform to the brands listed in ftyp
type Box struct {
	base.Box
	MajorBrand       [4]byte
	MinorVersion     uint32
	CompatibleBrands [][4]byte
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, [4]byte{}, 0, [][4]byte{}}
}

func (Box) Type() string {
	return FTYP
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s, majorbrand=%s, minorversion=%d, compatiblebrands=%s",
		b.Info(),
		b.MajorBrand,
		b.MinorVersion,
		b.CompatibleBrands,
	)
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	copy(b.MajorBrand[:], src[:4])
	b.MinorVersion = binary.BigEndian.Uint32(src[4:8])

	dst := [4]byte{}
	for i := 8; i < len(src); i += 4 {
		copy(dst[:], src[i:i+4])
		b.CompatibleBrands = append(b.CompatibleBrands, dst)
	}

	return len(src), nil
}
