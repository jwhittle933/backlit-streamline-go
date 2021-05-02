// Package dinf (Data Information)
package dinf

const (
	DINF string = "dinf"
)

// Box is ISOBMFF dinf box type
type Box struct {
	//
}

func (Box) Type() string {
	return DINF
}