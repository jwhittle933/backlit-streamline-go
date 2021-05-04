package esds

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsd/mp4a/esds/descriptor"
)

const (
	ESDS                       string = "esds"
	DescriptorTag                     = 0x03
	DecoderConfigDescriptorTag        = 0x04
	DecSpecificInfoTag                = 0x05
	SLConfigDescriptorTag             = 0x06
)

// Box is ES Descriptor Box
type Box struct {
	base.Box
	Descriptors []descriptor.Descriptor
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, make([]descriptor.Descriptor, 0)}
}

func (Box) Type() string {
	return ESDS
}

func (b *Box) String() string {
	return b.Info().String() + ", status=\033[35mINCOMPLETE\033[0m"
}

func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
