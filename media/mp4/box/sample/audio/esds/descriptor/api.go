package descriptor

import (
	"github.com/jwhittle933/streamline/media/mp4/box/sample/audio/esds/descriptor/decoder"
	"github.com/jwhittle933/streamline/media/mp4/box/sample/audio/esds/descriptor/es"
)

const (
	Tag                        = 0x03
	DecoderConfigDescriptorTag = 0x04
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
	case Tag:
		return name == "EsDescriptor"
	case DecoderConfigDescriptorTag:
		return name == "DecoderConfigDescriptor"
	}

	return name == "Data"
}
