// Package pssh (Protection System Specific Header)
package pssh

import (
	"fmt"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	PSSH string = "pssh"
)

// Box is ISO BMFF pssh box type
type Box struct {
	base.Box
	SystemID [16]byte
	KIDCount uint32
	KIDs     []KID
	DataSize int32
	Data     []byte
}

type KID [16]byte

func New(i *box.Info) box.Boxed {
	return &Box{
		base.Box{BoxInfo: i},
		[16]byte{},
		0,
		make([]KID, 0),
		0,
		make([]byte, 0),
	}
}

func (Box) Type() string {
	return PSSH
}

func (b Box) String() string {
	return fmt.Sprintf(
		"[%s] hex=%s, offset=%d, size=%d, header=%d",
		b.BoxInfo.Type.String(),
		b.BoxInfo.Type.HexString(),
		b.BoxInfo.Offset,
		b.BoxInfo.Size,
		b.BoxInfo.HeaderSize,
	)
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	// iteratively parse children
	return len(src), nil
}
