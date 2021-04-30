package styp

import "github.com/jwhittle933/streamline/pkg/media/mp4/box"

const (
	STYP string = "styp"
)

type Box struct {
	BoxInfo *box.Info
	MajorBrand [4]byte
	MinorVersion uint32
	CompatibleBrands [][4]byte
}

func New(i *box.Info) box.Boxed {
	return &Box{BoxInfo: i}
}

func (Box) Type() string {
	return STYP
}

func (b *Box) Info() *box.Info {
	return b.BoxInfo
}