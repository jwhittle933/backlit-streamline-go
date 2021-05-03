package gnre

const (
	GNRE string = "gnre"
)

type Box struct {
	//
}

func (Box) Type() string {
	return GNRE
}

