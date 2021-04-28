package meta

const (
	META string = "meta"
)

// Box is ISOBMFF meta box type
type Box struct {
	//
}

func (Box) Type() string {
	return META
}

