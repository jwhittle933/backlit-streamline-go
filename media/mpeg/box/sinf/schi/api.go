// Package schi (Scheme Information)
package schi

import (
	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/box/sinf/schi/tenc"
	"github.com/jwhittle933/streamline/media/mpeg/children"
)

const (
	SCHI string = "schi"
)

var (
	Children = children.Registry{
		tenc.TENC: tenc.New,
	}
)

type Box struct {
	base.Box
	SchemeType    [4]byte
	SchemeVersion uint32
	SchemeUri     []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, [4]byte{}, 0, make([]byte, 0)}
}

func (Box) Type() string {
	return SCHI
}

func (b *Box) String() string {
	return b.Info().String()
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
