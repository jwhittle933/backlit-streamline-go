package dinf

const (
	DINF string = "dinf"
)

// Dinf is ISOBMFF dinf box type
type Dinf struct {
	//
}

func (d Dinf) Type() string {
	return DINF
}