package unknown

import (
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	UNKNOWN = "unkn"
)

type Box struct {
	base2.Box
	Data []byte
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, []byte{}}
}

func (Box) Type() string {
	return UNKNOWN
}

func (b *Box) Info() *box2.Info {
	return b.BoxInfo
}

func (b Box) String() string {
	return b.Info().String()
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
