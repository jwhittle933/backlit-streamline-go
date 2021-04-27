package mdat

const (
	MDAT string = "mdat"
)

// Box is ISOBMFF mdat box type
type Box struct {
	Data []byte
}

func (Box) Type() string {
	return MDAT
}
