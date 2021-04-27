package urn

const (
	URN string = "urn "
	SelfContained uint = 0x000001
)

type Urn struct {
	Name     string
	Location string
}

func (u Urn) Type() string {
	return URN
}
