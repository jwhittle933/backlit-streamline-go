package styp

const (
	STYP string = "styp"
)

type Box struct {
	MajorBrand [4]byte
	MinorVersion uint32
	CompatibleBrands [][4]byte
}

func (Box) Type() string {
	return STYP
}