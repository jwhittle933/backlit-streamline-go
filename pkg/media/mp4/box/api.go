package box

import (
	"fmt"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/boxtype"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov"
	"io"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box/ftyp"
)

const (
	SmallHeader = 8
	LargeHeader = 16
)

var BoxRegistry = map[string]Boxed{
	"ftyp": ftyp.Box{},
	"moov": moov.Box{},
}

type Boxed interface {
	//io.Writer
	Type() string
	//Version() uint8
	//Children() []Boxed
}

type Box struct {
	Info *Info
}

func New() *Box {
	return &Box{}
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
		"[%s] hex=%s, offset=%d, size=%d, header=%d]",
		string(i.Type[:]),
		i.Type.String(),
		i.Offset,
		i.Size,
		i.HeaderSize,
	)
}
