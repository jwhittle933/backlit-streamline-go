// Package traf (Track Fragment)
package traf

import (
	"github.com/jwhittle933/streamline/media/mp4/box/children"
	"github.com/jwhittle933/streamline/media/mp4/box/subs"
)

const (
	TRAF string = "traf"
)

var (
	Children = children.Registry{
		subs.SUBS: subs.New,
	}
)

type Box struct {
	//
}

func (Box) Type() string {
	return TRAF
}
