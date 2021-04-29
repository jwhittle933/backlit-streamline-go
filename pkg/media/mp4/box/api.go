package box

import "io"

const (
	SmallHeader = 8
	LargeHeader = 16
)

type Boxed interface {
	io.Writer
	Type() string
	Version() uint8
	Children() []Boxed
}

type Info struct {
	Offset      uint64
	Size        uint64
	Type        [4]byte
	HeaderSize  uint64
	ExtendToEOF bool
}

func (i *Info) SeekPayload(s io.Seeker) (int64, error) {
	return s.Seek(int64(i.Offset+i.HeaderSize), io.SeekStart)
}
