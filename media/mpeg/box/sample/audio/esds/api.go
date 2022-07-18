package esds

import (
	"errors"
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/fullbox"
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
	fullbox.Box
	EsDescrTag            byte
	EsID                  uint16
	FlagsAndPriority      byte
	DecoderConfigDescrTag byte
	ObjectType            byte
	StreamType            byte
	BufferSizeDB          uint32
	MaxBitrate            uint32
	AvgBitrate            uint32
	DecSpecificInfoTag    byte
	DecConfig             []byte
	SLConfigDescrTag      byte
	SLConfigValue         byte
	nrExtraSizeBytes      int
}

// New returns a new zeroed Box
func New(i *box.Info) box.Boxed {
	return &Box{
		*fullbox.New(i),
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		make([]byte, 0),
		0,
		0,
		0,
	}
}

func (Box) Type() string {
	return ESDS
}

func (b *Box) String() string {
	return fmt.Sprintf(
		"%s, version=%d, flags=%d, esdesctag=%d, esid=%d, flags_priority=%d, decoder_config_desc_tag=%d, objecttype=%d, streamtype=%d, buffersizedb=%d, maxbitrate=%d, avgbitrate=%d, decconfig=%+v, slconfigdesctag=%d, slconfig=%d, extrabytes=%d",
		b.Info(),
		b.Version,
		b.Flags,
		b.EsDescrTag,
		b.EsID,
		b.FlagsAndPriority,
		b.DecoderConfigDescrTag,
		b.ObjectType,
		b.StreamType,
		b.BufferSizeDB,
		b.MaxBitrate,
		b.AvgBitrate,
		b.DecConfig,
		b.SLConfigDescrTag,
		b.SLConfigValue,
		b.nrExtraSizeBytes,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)

	b.EsDescrTag = sr.Uint8()

	_, bytesRead := b.sizeOfInstance(sr)
	b.nrExtraSizeBytes += bytesRead - 1

	b.EsID = sr.Uint16()
	b.FlagsAndPriority = sr.Uint8()
	b.DecoderConfigDescrTag = sr.Uint8()

	_, bytesRead = b.sizeOfInstance(sr)
	b.nrExtraSizeBytes += bytesRead - 1

	b.ObjectType = sr.Uint8()

	streamTypeAndBufferSizeDB := sr.Uint32()
	b.StreamType = byte(streamTypeAndBufferSizeDB >> 24)
	b.BufferSizeDB = streamTypeAndBufferSizeDB & 0xffffff
	b.MaxBitrate = sr.Uint32()
	b.AvgBitrate = sr.Uint32()
	b.DecSpecificInfoTag = sr.Uint8()
	size, bytesRead := b.sizeOfInstance(sr)
	b.nrExtraSizeBytes += bytesRead
	b.DecConfig = sr.Slice(size)
	b.SLConfigDescrTag = sr.Uint8()

	size, bytesRead = b.sizeOfInstance(sr)
	b.nrExtraSizeBytes += bytesRead
	if size != 1 {
		return len(src), errors.New("cannot handle SlConfigDescr not equal to 1 byte")
	}

	b.SLConfigValue = sr.Uint8()

	return box.FullRead(len(src))
}

func (b Box) sizeOfInstance(sr *slicereader.Reader) (int, int) {
	tmp := sr.Uint8()
	nrBytesRead := 1
	sizeOfInstance := int(tmp & 0x7f)

	for {
		if (tmp >> 7) == 0 {
			break
		}

		tmp = sr.Uint8()
		nrBytesRead++
		sizeOfInstance = sizeOfInstance<<7 | int(tmp&0x7f)
	}

	return sizeOfInstance, nrBytesRead
}
