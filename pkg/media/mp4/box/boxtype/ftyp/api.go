package ftyp

const (
	FTYP string = "ftyp"
)

// Box is ISOBMFF ftyp box type
type Box struct {
	MajorBrand       [4]byte
	MinorVersion     uint32
	CompatibleBrands [][4]byte
}

func (Box) Type() string {
	return FTYP
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
