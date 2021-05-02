// Package smhd (Sound Media Header)
package smhd

const (
	SMHD string = "smhd"
)

type Box struct {
	Balance   int16
	_reserved uint16
}

func (Box) Type() string {
	return SMHD
}
