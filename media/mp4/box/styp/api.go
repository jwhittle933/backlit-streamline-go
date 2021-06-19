package styp

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	STYP string = "styp"
)

type Box struct {
	base.Box
	MajorBrand       [4]byte
	MinorVersion     uint32
	CompatibleBrands [][4]byte
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, [4]byte{}, 0, make([][4]byte, 0)}
}

func (Box) Type() string {
	return STYP
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s",
		b.Info(),
	)
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	copy(b.MajorBrand[:], sr.Slice(4))
	b.MinorVersion = sr.Uint32()

	if sr.Length() > 8 {
		b.CompatibleBrands = make([][4]byte, sr.Length()/4)
		for i := 0; i < sr.Length(); i += 4 {
			cb := [4]byte{}
			copy(cb[:], sr.Slice(4))
			b.CompatibleBrands[i] = cb
		}
	}

	return box.FullRead(len(src))
}
