// Package saio (Sample Auxiliary Information Offsets)
package saio

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/fullbox"
)

const (
	SAIO string = "saio"
)

type Box struct {
	fullbox.Box
	AuxInfoType          [4]byte // string
	AuxInfoTypeParameter uint32
	EntryCount           uint32
	Offset               []int64
}

func (Box) Type() string {
	return SAIO
}

func (b Box) String() string {
	return fmt.Sprintf("%s", b.Info())
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)
	if b.Flags&0x01 != 0 {
		copy(b.AuxInfoType[:], sr.Slice(4))
		b.AuxInfoTypeParameter = sr.Uint32()
	}

	b.EntryCount = sr.Uint32()
	b.Offset = make([]int64, b.EntryCount)

	if b.Version == 0 {
		for i := 0; i < len(b.Offset); i++ {
			b.Offset[i] = int64(sr.Uint32())
		}
	} else {
		for i := 0; i < len(b.Offset); i++ {
			b.Offset[i] = int64(sr.Uint64())
		}
	}

	return box.FullRead(len(src))
}
