package mfra

const (
	MFRA string = "mfra"
)

type Box struct {
	//
}

func (Box) Type() string {
	return MFRA
}
