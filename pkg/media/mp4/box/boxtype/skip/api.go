package skip

const (
	SKIP string = "skip"
)

type Box struct {
	Data []uint8
}

func (Box) Type() string {
	return SKIP
}
