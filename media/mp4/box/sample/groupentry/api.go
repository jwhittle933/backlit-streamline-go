package groupentry

import (
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/boxtype"
)

// Sample Behaves like a box but has no header
type Sample interface {
	box.Boxed
}

func ScanEntry(name string, src []byte, factoryFn box.Factory) (Sample, error) {
	b := factoryFn(&box.Info{
		Type: boxtype.New([4]byte{name[0], name[1], name[2], name[3]}),
		Size: uint64(len(src)),
	})

	if _, err := b.Write(src); err != nil {
		return b, err
	}

	return b, nil
}

