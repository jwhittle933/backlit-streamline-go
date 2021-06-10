// Package box defines functions to read from
// and write to Boxes within in ISO BMFF box/atom
package box

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	boxtype2 "github.com/jwhittle933/streamline/media/mp4/box/boxtype"
	"io"
)

const (
	SmallHeader uint64 = 8
	LargeHeader uint64 = 16
)

type Factory func(*Info) Boxed

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

type Info struct {
	Offset      uint64
	Size        uint64
	Type        boxtype2.BoxType
	HeaderSize  uint64
	ExtendToEOF bool
}

func (i Info) String() string {
	return fmt.Sprintf(
		"[%s] offset=%d, size=%d, header=%d",
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

func SeekPayload(s io.Seeker, info *Info) (int64, error) {
	return s.Seek(int64(info.Offset+info.HeaderSize), io.SeekStart)
}

func ScanInfo(r io.ReadSeeker, i *Info) error {
	off, err := r.Seek(0, io.SeekCurrent)
	if err != nil {
		return err
	}

	i.Offset = uint64(off)
	i.HeaderSize = SmallHeader

	buf := bytes.NewBuffer(make([]byte, 0, i.HeaderSize))
	if _, err := io.CopyN(buf, r, int64(i.HeaderSize)); err != nil {
		return err
	}

	data := buf.Bytes()
	i.Size = uint64(binary.BigEndian.Uint32(data))
	i.Type = boxtype2.New([4]byte{data[4], data[5], data[6], data[7]})

	if i.Size == 0 {
		off, _ = r.Seek(0, io.SeekEnd)
		i.Size = uint64(off) - i.Offset
		i.ExtendToEOF = true
		if _, err := SeekPayload(r, i); err != nil {
			return err
		}

		return nil
	}

	if i.Size == 1 {
		buf.Reset()
		if _, err := io.CopyN(buf, r, 8); err != nil {
			return err
		}

		i.HeaderSize = LargeHeader
		i.Size = binary.BigEndian.Uint64(buf.Bytes())

		return nil
	}

	return nil
}

func FullRead(read int) (int, error) {
	return read, nil
}
