package box

import (
	"fmt"
	"io"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box/boxtype"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/ftyp"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/mdat"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moof"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/styp"
)

const (
	SmallHeader = 8
	LargeHeader = 16
)

var BoxRegistry = map[string]Boxed{
	"ftyp": ftyp.Box{},
	"mdat": mdat.Box{},
	"moov": moov.Box{},
	"moof": moof.Box{},
	"styp": styp.Box{},
	"free": mdat.Box{},
}

type Boxed interface {
	//io.Writer
	Type() string
	//Version() uint8
	//Children() []Boxed
}

type Box struct {
	Boxed
	Info *Info
}

func New(i *Info) *Box {
	return &Box{Info: i}
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
		string(i.Type[:]),
		i.Type.String(),
		i.Offset,
		i.Size,
		i.HeaderSize,
	)
}
