// Package stbl (Sample Table)
package stbl

const (
	STBL string = "stbl"
)

type Box struct {
	//
}

func (Box) Type() string {
	return STBL
}