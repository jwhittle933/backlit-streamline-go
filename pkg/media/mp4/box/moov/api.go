package moov

import (
	"bytes"
	"fmt"
	"io"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/children"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/free"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/coin"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/meta"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/mvex"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/mvhd"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/pssh"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/unknown"
)

const (
	MOOV string = "moov"
)

var Children = children.Registry{
	coin.COIN: coin.New,
	free.FREE: free.New,
	meta.META: meta.New,
	mvex.MVEX: mvex.New,
	mvhd.MVHD: mvhd.New,
	pssh.PSSH: pssh.New,
	trak.TRAK: trak.New,
}

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
	s := fmt.Sprintf("%s, boxes=%d\n", b.Info().String(), len(b.Children))

	for _, c := range b.Children {
		s += fmt.Sprintf("  %s\n", c.String())
	}

	return s
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	r := bytes.NewReader(src)

	for {
		child, err := b.scan(r)
		if err == io.EOF {
			return len(src), nil
		}

		if err != nil || child == nil {
			break
		}

		b.Children = append(b.Children, child)
	}

	return len(src), nil
}

func (b *Box) scan(r io.ReadSeeker) (box.Boxed, error) {
	i := &box.Info{}
	err := box.ScanInfo(r, i)

	if err != nil {
		return nil, err
	}

	var boxFactory children.BoxFactory
	var found bool
	if boxFactory, found = Children[i.Type.String()]; !found {
		boxFactory = unknown.New
	}

	child := boxFactory(i)

	if _, err := io.CopyN(child, r, int64(i.Size-i.HeaderSize)); err != nil {
		return nil, err
	}

	return child, nil
}
