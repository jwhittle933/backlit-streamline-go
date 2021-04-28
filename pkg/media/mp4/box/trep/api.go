package trep

const (
	TREP string = "trep"
)

type Box struct {
	TrackID uint32
}

func (Box) Type() string {
	return TREP
}
