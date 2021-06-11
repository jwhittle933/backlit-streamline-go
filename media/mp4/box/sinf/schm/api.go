// Package schm (Scheme Type)
package schm

import (
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	SCHM string = "schm"
)

type Box struct {
	base2.Box
	SchemeType    [4]byte
	SchemeVersion uint32
	SchemeUri     []byte
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, [4]byte{}, 0, make([]byte, 0)}
}

func (Box) Type() string {
	return SCHM
}

func (b *Box) String() string {
	return b.Info().String()
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}