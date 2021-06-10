// Package mfhd (Movie Fragment Header)
package mfhd

const (
	MFHD string = "mfhd"
)

type Box struct {
	SequenceNumber uint32
}

func (Box) Type() string {
	return MFHD
}
