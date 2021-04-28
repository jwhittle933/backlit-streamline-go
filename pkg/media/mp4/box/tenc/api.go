package tenc

const (
	TENC string = "tenc"
)

type Box struct {
	_reserved              uint8
	DefaultCryptByteBlock  uint8
	DefaultSkipByteBlock   uint8
	DefaultIsProtected     uint8
	DefaultPerSampleIVSize uint8
	DefaultKID             [16]byte
	DefaultConstantIVSize  uint8
	DefaultConstantIV      []byte
}

func (Box) Type() string {
	return TENC
}
