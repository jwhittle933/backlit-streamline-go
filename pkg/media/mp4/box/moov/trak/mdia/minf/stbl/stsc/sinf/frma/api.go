// Package frma (Original Format)
package frma

const (
	FRMA string = "frma"
)

// Box is ISOBMFF frma box type
type Box struct {
	DataFormat [4]byte
}

func (Box) Type() string {
	return FRMA
}
