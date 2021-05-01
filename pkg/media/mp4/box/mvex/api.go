// Package mvex for ISO BMFF Movie Extends
package mvex

const (
	MVEX string = "mvex"
)

type Box struct {
	//
}

func (Box) Type() string {
	return MVEX
}