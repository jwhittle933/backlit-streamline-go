package edts

const (
	EDTS string = "edts"
)

// Edts is ISOBMFF edts box type
type Edts struct {
	//
}

func (Edts) Type() string {
	return EDTS
}
