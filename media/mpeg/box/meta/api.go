// Package meta for multi-track required meta-data
package meta

import "github.com/jwhittle933/streamline/media/mpeg/fullbox"

const (
	META string = "meta"
)

// Box is ISOBMFF meta box type
type Box struct {
	fullbox.Box
}

func (Box) Type() string {
	return META
}
