// Package traf (Track Fragment)
package traf

const (
	TRAF string = "traf"
)


type Box struct {
	//
}

func (Box) Type() string {
	return TRAF
}