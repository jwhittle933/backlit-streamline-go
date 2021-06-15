// Package mfra (Movie Fragment Random Access)
package mfra

import (
	"fmt"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
	"github.com/jwhittle933/streamline/media/mp4/box/children"
)

const (
	MFRA string = "mfra"
)

var (
	Children children.Registry
)

type Box struct {
	base.Box
	Children []box.Boxed
}

func (Box) Type() string {
	return MFRA
}

func (b Box) String() string {
	s := fmt.Sprintf("%s", b.Info())

	for _, c := range b.Children {
		s += fmt.Sprintf("\n  %s", c)
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	return box.FullRead(len(src))
}
