package moov

import (
	"bytes"
	"fmt"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	children2 "github.com/jwhittle933/streamline/media/mp4/box/children"
	free2 "github.com/jwhittle933/streamline/media/mp4/box/free"
	coin2 "github.com/jwhittle933/streamline/media/mp4/box/moov/coin"
	meta2 "github.com/jwhittle933/streamline/media/mp4/box/moov/meta"
	mvex2 "github.com/jwhittle933/streamline/media/mp4/box/moov/mvex"
	mvhd2 "github.com/jwhittle933/streamline/media/mp4/box/moov/mvhd"
	pssh2 "github.com/jwhittle933/streamline/media/mp4/box/moov/pssh"
	trak2 "github.com/jwhittle933/streamline/media/mp4/box/moov/trak"
	scanner2 "github.com/jwhittle933/streamline/media/mp4/box/scanner"
)

const (
	MOOV string = "moov"
)

var Children = children2.Registry{
	coin2.COIN: coin2.New,
	free2.FREE: free2.New,
	meta2.META: meta2.New,
	mvex2.MVEX: mvex2.New,
	mvhd2.MVHD: mvhd2.New,
	pssh2.PSSH: pssh2.New,
	trak2.TRAK: trak2.New,
}

// Box is ISO BMFF moov box.
// The box contains no data itself
type Box struct {
	base2.Box
	Children []box2.Boxed
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, make([]box2.Boxed, 0)}
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
	s := scanner2.New(bytes.NewReader(src))

	var err error
	b.Children, err = s.ScanAllChildren(Children)
	if err != nil {
		return 0, err
	}

	return box2.FullRead(len(src))
}
