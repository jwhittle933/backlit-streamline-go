package edts

const (
	EDTS string = "edts"
)

// Box is ISOBMFF edts box type
type Box struct {
	//
}

func (Box) Type() string {
	return EDTS
}
