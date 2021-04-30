// Package ftyp
// See: http://www.ftyps.com/what.html#:~:text=It%20only%20pertains%20to%20MP4%20or%20newer%20QuickTime,QuickTime%20terminology%29%20or%20box%20type%20%28in%20MP4%20terminology%29.
package ftyp

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
)

const (
	FTYP string = "ftyp"
)

// Box is ISOBMFF ftyp box type
// If the segment type box (styp) is not present
// the segment must conform to the brands listed in ftyp
type Box struct {
	BoxInfo          *box.Info
	MajorBrand       [4]byte
	MinorVersion     uint32
	CompatibleBrands [][4]byte
}

func New(i *box.Info) box.Boxed {
	return Box{BoxInfo: i}
}

func (Box) Type() string {
	return FTYP
}

func (b Box) Info() *box.Info {
	return b.BoxInfo
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	return 0, nil
}

func (b *Box) AddCompatibleBrand(cb [4]byte) bool {
	if !b.HasCompatibleBrand(cb) {
		b.CompatibleBrands = append(b.CompatibleBrands, cb)
		return true
	}

	return false
}

func (b *Box) RemoveCompatibleBrand(cb [4]byte) bool {
	for i := 0; i < len(b.CompatibleBrands); i++ {
		if b.CompatibleBrands[i] != cb {
			continue
		}

		b.CompatibleBrands = append(b.CompatibleBrands[:i], b.CompatibleBrands[i:]...)
		return true
	}

	return false
}

func (b *Box) HasCompatibleBrand(cb [4]byte) bool {
	for i := range b.CompatibleBrands {
		if b.CompatibleBrands[i] == cb {
			return true
		}
	}

	return false
}
