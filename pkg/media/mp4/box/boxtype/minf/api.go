package minf

const (
	MINF string = "minf"
)

type Box struct {
	//
}

func (Box) Type() string {
	return MINF
}
