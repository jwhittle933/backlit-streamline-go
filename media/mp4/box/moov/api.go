package moov

import (
	"bytes"
	"fmt"

	"github.com/jwhittle933/streamline/media/mp4/base"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/free"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/coin"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/meta"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/mvex"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/mvhd"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/pssh"
	"github.com/jwhittle933/streamline/media/mp4/box/moov/trak"
	"github.com/jwhittle933/streamline/media/mp4/box/scanner"
	"github.com/jwhittle933/streamline/media/mp4/box/udta"
	"github.com/jwhittle933/streamline/media/mp4/children"
)

const (
	MOOV string = "moov"
)

var (
	Children = children.Registry{
		coin.COIN: coin.New,
		free.FREE: free.New,
		meta.META: meta.New,
		mvex.MVEX: mvex.New,
		mvhd.MVHD: mvhd.New,
		pssh.PSSH: pssh.New,
		trak.TRAK: trak.New,
		udta.UDTA: udta.New,
	}
)

// Box is ISO BMFF moov box.
// The box contains no data itself
type Box struct {
	base.Box
	Children []box.Boxed
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, make([]box.Boxed, 0)}
}

func (Box) Type() string {
	return MOOV
}

func (b Box) String() string {
	s := fmt.Sprintf("%s, children=%d", b.Info(), len(b.Children))

	for _, c := range b.Children {
		s += fmt.Sprintf("\n  %s", c.String())
	}

	return s
}

// Write writes to *Box from `src`
// Satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	s := scanner.New(bytes.NewReader(src))

	var err error
	b.Children, err = s.ScanAllChildren(Children)
	if err != nil {
		return 0, err
	}

	return box.FullRead(len(src))
}
