package sidx

import (
	"fmt"
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	SIDX string = "sidx"
)

/*
Definition according to ISO/IEC 14496-12 Section 8.16.3.2
aligned(8) class SegmentIndexBox extends FullBox(‘sidx’, version, 0) {
	unsigned int(32) reference_ID;
	unsigned int(32) timescale;
	if (version==0) {
		unsigned int(32) earliest_presentation_time;
		unsigned int(32) first_offset;
	} else {
		unsigned int(64) earliest_presentation_time; unsigned int(64) first_offset;
	}
	unsigned int(16) reserved = 0;
	unsigned int(16) reference_count;
	for(i=1; i <= reference_count; i++) {
		bit (1)           reference_type;
		unsigned int(31)  referenced_size;
		unsigned int(32)  subsegment_duration;
		bit(1)            starts_with_SAP;
		unsigned int(3)   SAP_type;
		unsigned int(28)  SAP_delta_time;
    }
}
*/

type Box struct {
	base.Box
	Version                  byte
	Flags                    uint32
	ReferenceID              uint32
	Timescale                uint32
	EarliestPresentationTime uint64
	FirstOffset              uint64
	References               []Reference
}

type Reference struct {
	ReferencedSize     uint32
	SubsegmentDuration uint32
	SAPDeltaTime       uint32
	ReferenceType      uint8 // 1 bit
	StartsWithSAP      uint8 // 1 bit
	SAPType            uint8
}

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		0,
		0,
		0,
		0,
		0,
		0,
		make([]Reference, 0),
	}
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s, version=%d, flags=%d, referenceid=%d timescale=%d, earliest_presentation_time=%d, firstoffset=%d, references=%d",
		b.Info(),
		b.Version,
		b.Flags,
		b.ReferenceID,
		b.Timescale,
		b.EarliestPresentationTime,
		b.FirstOffset,
		len(b.References),
	)
}

func (Box) Type() string {
	return SIDX
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	versionAndFlags := sr.Uint32()
	b.Version = byte(versionAndFlags >> 24)
	b.Flags = versionAndFlags & box.FlagsMask

	b.ReferenceID = sr.Uint32()
	b.Timescale = sr.Uint32()

	if b.Version == 0 {
		b.EarliestPresentationTime = uint64(sr.Uint32())
		b.FirstOffset = uint64(sr.Uint32())
	} else {
		b.EarliestPresentationTime = sr.Uint64()
		b.FirstOffset = sr.Uint64()
	}

	sr.Skip(2)
	for i, refCount := 0, sr.Uint16(); i < int(refCount); i++ {
		work := sr.Uint32()
		ref := Reference{
			ReferenceType:      uint8(work >> 31),
			ReferencedSize:     work & 0x7fffffff,
			SubsegmentDuration: sr.Uint32(),
		}

		work = sr.Uint32()
		ref.StartsWithSAP = uint8(work >> 31)
		ref.SAPType = uint8((work >> 28) & 0x07)
		ref.SAPDeltaTime = work & 0x0fffffff

		b.References = append(b.References, ref)
	}

	return box.FullRead(len(src))
}
