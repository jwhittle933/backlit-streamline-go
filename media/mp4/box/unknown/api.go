package unknown

import (
	"fmt"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	UNKNOWN = "unkn"
)

type Box struct {
	base.Box
	Data []byte
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, []byte{}}
}

func (Box) Type() string {
	return UNKNOWN
}

func (b *Box) Info() *box.Info {
	return b.BoxInfo
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s \033[0;31mUNKNOWN\033[0m",
		b.Info(),
	)
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
