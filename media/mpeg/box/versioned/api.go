package versioned

type Box struct {
	Version uint8
}

func (b *Box) WriteVersion(src []byte) []byte {
	b.Version = src[0]
	return src[1:]
}
