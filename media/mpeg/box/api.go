// Package box defines functions to read from
// and write to Boxes within in ISO BMFF box/atom
package box

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/jwhittle933/streamline/media/mpeg/boxtype"
)

const (
	SmallHeader uint64 = 8
	LargeHeader uint64 = 16
	FlagsMask          = 0x00ffffff
)

type Factory func(*Info) Boxed

type Typed interface {
	Type() string
}

type Informed interface {
	Info() *Info
}

type Finder interface {
	Find(name string) Boxed
}

type Boxed interface {
	io.Writer
	fmt.Stringer
	Typed
	Informed
}

type Info struct {
	Offset      uint64
	Size        uint64
	Type        boxtype.BoxType
	HeaderSize  uint64
	ExtendToEOF bool
}

func (i Info) String() string {
	return fmt.Sprintf(
		"[\033[1;36m%s\033[0m] offset=%d, size=%d, header=%d",
		i.Type.String(),
		i.Offset,
		i.Size,
		i.HeaderSize,
	)
}

func (i Info) Read(dst []byte) (int, error) {
	if i.Size > 1<<32 {
		return 0, errors.New("header too large")
	}

	binary.BigEndian.PutUint32(dst, uint32(i.Size))
	copy(dst, i.Type[:])

	return int(i.HeaderSize), nil
}

// FullRead sugar function for box write returns
func FullRead(read int) (int, error) {
	return read, nil
}
