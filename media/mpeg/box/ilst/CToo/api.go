package CToo

import (
	"bytes"
	"fmt"
	"github.com/jwhittle933/streamline/media/mpeg/children"

	"github.com/jwhittle933/streamline/media/mpeg/base"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/box/data"
	"github.com/jwhittle933/streamline/media/mpeg/box/scanner"
)

const (
	CTOO = "\xa9too"
)

type Box struct {
	base.Box
	Children []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, make([]box.Boxed, 0)}
}

func (Box) Type() string {
	return CTOO
}

func (b Box) String() string {
	s := fmt.Sprintf("%s", b.Info())

	for _, c := range b.Children {
		s += fmt.Sprintf("\n          %s", c)
	}

	return s
}

func (b *Box) Write(src []byte) (int, error) {
	s := scanner.New(bytes.NewReader(src))

	var err error
	b.Children, err = s.ScanAllChildren(children.Registry{data.DATA: data.New})
	if err != nil {
		return 0, err
	}
	return box.FullRead(len(src))
}
