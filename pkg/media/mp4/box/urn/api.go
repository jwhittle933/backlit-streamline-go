package urn

const (
	URN string = "urn "
	SelfContained uint = 0x000001
)

type Box struct {
	Name     string
	Location string
}

func (Box) Type() string {
	return URN
}
