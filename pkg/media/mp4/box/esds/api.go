package esds

import "github.com/jwhittle933/streamline/pkg/media/mp4/box/esds/descriptor"

const (
	ESDS                       string = "esds"
	EsdsDescriptorTag                 = 0x03
	DecoderConfigDescriptorTag        = 0x04
	DecSpecificInfoTag                = 0x05
	SLConfigDescriptorTag             = 0x06
)

// Box is ES Descriptor Box
type Box struct {
	Descriptors []descriptor.Descriptor
}

func (Box) Type() string {
	return ESDS
}
