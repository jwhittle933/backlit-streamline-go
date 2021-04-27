package free

const (
	FREE string = "free"
)

type Box struct {
	Data []uint8
}

func (Box) Type() string {
	return FREE
}
