package descriptor

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/esds/descriptor/decoder"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/esds/descriptor/es"
)

const (
	EsdsDescriptorTag                 = 0x03
	DecoderConfigDescriptorTag        = 0x04
)

type Descriptor struct {
	Tag                     int8
	Size                    uint32
	ESDescriptor            *es.Descriptor
	DecoderConfigDescriptor *decoder.ConfigDescriptor
	Data                    []byte
}

func (d *Descriptor) IsOptFieldEnabled(name string) bool {
	switch d.Tag {
	case EsdsDescriptorTag:
		return name == "EsDescriptor"
	case DecoderConfigDescriptorTag:
		return name == "DecoderConfigDescriptor"
	}

	return name == "Data"
}
