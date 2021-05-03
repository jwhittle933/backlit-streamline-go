package cprt

const (
	CPRT string = "cprt"
)

type Box struct {
	//
}

func (Box) Type() string {
	return CPRT
}

