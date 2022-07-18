package enca

const (
	ENCA string = "enca"
)

type Box struct{}

func (Box) Type() string {
	return ENCA
}
