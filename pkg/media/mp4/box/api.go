package box

import (
	"fmt"
	"io"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box/boxtype"
)

const (
	SmallHeader = 8
	LargeHeader = 16
)

type Typed interface {
	Type() string
}

type Informed interface {
	Info() *Info
}

type Boxed interface {
	io.Writer
	fmt.Stringer
	Typed
	Informed
}

type Box struct {
	Boxed
	BoxInfo *Info
}

func New(i *Info) *Box {
	return &Box{BoxInfo: i}
}

func (b Box) Type() string {
	return string(b.BoxInfo.Type[:])
}

func (b Box) Info() *Info {
	return b.BoxInfo
}

type Info struct {
	Offset      uint64
	Size        uint64
	Type        boxtype.BoxType
	HeaderSize  uint64
	ExtendToEOF bool
}

func (i *Info) SeekPayload(s io.Seeker) (int64, error) {
	return s.Seek(int64(i.Offset+i.HeaderSize), io.SeekStart)
}

func (i *Info) String() string {
	return fmt.Sprintf(
		"[%s] hex=%s, offset=%d, size=%d, header=%d",
		string(i.Type.String()),
		i.Type.HexString(),
		i.Offset,
		i.Size,
		i.HeaderSize,
	)
}
