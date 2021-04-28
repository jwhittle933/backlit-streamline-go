package trak

const (
	TRAK string = "trak"
)

type Box struct {
	//
}

func (Box) Type() string {
	return TRAK
}
