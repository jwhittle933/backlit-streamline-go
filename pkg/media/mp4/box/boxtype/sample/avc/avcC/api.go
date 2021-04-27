package avcC

const (
	AVCC string = "avcC"
)

type Box struct {
	//
}

func (Box) Type() string {
	return AVCC
}
