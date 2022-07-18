package wave

const (
	WAVE string = "wave"
)

type Box struct {
	//
}

func (Box) Type() string {
	return WAVE
}
