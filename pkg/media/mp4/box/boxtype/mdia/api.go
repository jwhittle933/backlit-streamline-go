package mdia

const (
	MDIA string = "mdia"
)

type Box struct {
	//
}

func (Box) Type() string {
	return MDIA
}
