package mfro

const (
	MFRO string = "mfro"
)

type Box struct {
	Size uint32
}

func (Box) Type() string {
	return MFRO
}