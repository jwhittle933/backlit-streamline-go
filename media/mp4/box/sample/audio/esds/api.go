package esds

import (
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
	descriptor2 "github.com/jwhittle933/streamline/media/mp4/box/sample/audio/esds/descriptor"
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
	base2.Box
	Descriptors []descriptor2.Descriptor
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, make([]descriptor2.Descriptor, 0)}
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
