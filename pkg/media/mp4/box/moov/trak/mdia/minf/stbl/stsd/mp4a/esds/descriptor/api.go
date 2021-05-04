package descriptor

import (
	decoder2 "github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsd/mp4a/esds/descriptor/decoder"
	es2 "github.com/jwhittle933/streamline/pkg/media/mp4/box/moov/trak/mdia/minf/stbl/stsd/mp4a/esds/descriptor/es"
)

const (
	Tag                        = 0x03
	DecoderConfigDescriptorTag = 0x04
)

type Descriptor struct {
	Tag                     int8
	Size                    uint32
	ESDescriptor            *es2.Descriptor
	DecoderConfigDescriptor *decoder2.ConfigDescriptor
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
