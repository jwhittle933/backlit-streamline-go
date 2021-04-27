package moof

const (
	MOOF string = "moof"
)

type Box struct {
	//
}

func (Box) Type() string {
	return MOOF
}
