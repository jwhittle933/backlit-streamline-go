package udta

import (
	"bytes"
	"fmt"

	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
	"github.com/jwhittle933/streamline/media/mp4/box/children"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/meta"
	"github.com/jwhittle933/streamline/media/mp4/box/scanner"
)

const (
	UDTA string = "udta"
)

var (
	Children = children.Registry{
		meta.META: meta.New,
	}
)

// Box is ISOBMFF udta box type
type Box struct {
	base.Box
	Children []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, make([]box.Boxed, 0)}
}

type GPPString struct {
	Pad      bool
	Language [3]byte
	Data     []byte
}

func (Box) Type() string {
	return UDTA
}

func (b Box) String() string {
	s := fmt.Sprintf("%s", b.Info())

	for _, c := range b.Children {
		s += fmt.Sprintf("\n    %s", c)
	}
	 return  s
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	s := scanner.New(bytes.NewReader(src))

	var err error
	b.Children, err = s.ScanAllChildren(Children)
	if err != nil {
		return 0, err
	}

	return box.FullRead(len(src))
}
