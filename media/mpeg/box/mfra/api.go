// Package mfra (Movie Fragment Random Access)
package mfra

import (
	"bytes"
	"fmt"

	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/box/mfra/mfro"
	"github.com/jwhittle933/streamline/media/mpeg/box/mfra/tfra"
	"github.com/jwhittle933/streamline/media/mpeg/box/scanner"
	"github.com/jwhittle933/streamline/media/mpeg/children"
)

const (
	MFRA string = "mfra"
)

var (
	Children = children.Registry{
		tfra.TFRA: tfra.New,
		mfro.MFRO: mfro.New,
	}
)

type Box struct {
	base.Box
	Children []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, make([]box.Boxed, 0)}
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
	s := scanner.New(bytes.NewReader(src))

	var err error
	b.Children, err = s.ScanAllChildren(Children)
	if err != nil {
		return 0, err
	}

	return box.FullRead(len(src))
}
