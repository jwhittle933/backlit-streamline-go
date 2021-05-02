// Package smhd (Subtitle Media Header)
package sthd

const (
	STHD string = "sthd"
)

type Box struct {
	Balance   int16
	_reserved uint16
}

func (Box) Type() string {
	return STHD
}
