package groupentry

import (
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/boxtype"
)

// Sample Behaves like a box but has no header
type Sample interface {
	box.Boxed
}

func ScanEntry(name string, src []byte, factoryFn box.Factory) (Sample, error) {
	b := factoryFn(&box.Info{
		Type: boxtype.FromString(name),
		Size: uint64(len(src)),
	})

	if _, err := b.Write(src); err != nil {
		return b, err
	}

	return b, nil
}
