package moov

const (
	MOOV string = "moov"
)

type Box struct {
	//
}

func (Box) Type() string {
	return MOOV
}
