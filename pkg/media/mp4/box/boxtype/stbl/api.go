package stbl

const (
	SMHD string = "smhd"
)

type Box struct {
	//
}

func (Box) Type() string {
	return SMHD
}