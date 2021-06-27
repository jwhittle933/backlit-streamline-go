package scanner

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"

	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/boxtype"
	"github.com/jwhittle933/streamline/media/mp4/children"
)

type Scanner interface {
	ScanAllChildren(knownChildren children.Registry) ([]box.Boxed, error)
	ScanFor(knownChildren children.Registry) (box.Boxed, error)
	ScanInfo(i *box.Info) error
	SeekPayload(info *box.Info) (int64, error)
}

type scanner struct {
	io.ReadSeeker
}

func New(r io.ReadSeeker) Scanner {
	return &scanner{r}
}

func (s *scanner) ScanAllChildren(knownChildren children.Registry) ([]box.Boxed, error) {
	found := make([]box.Boxed, 0)

	var child box.Boxed
	var err error
	for {
		child, err = s.ScanFor(knownChildren)
		if err == io.EOF {
			return found, nil
		}

		if err != nil || child == nil {
			break
		}

		found = append(found, child)
	}

	return found, err
}

func (s *scanner) ScanFor(knownChildren children.Registry) (box.Boxed, error) {
	i := &box.Info{}
	err := s.ScanInfo(i)

	if err != nil {
		return nil, err
	}

	factory := knownChildren.Get(i.Type.String())
	child := factory(i)

	if _, err := io.CopyN(child, s, int64(i.Size-i.HeaderSize)); err != nil {
		return child, err
	}

	return child, nil
}

func (s *scanner) ScanInfo(i *box.Info) error {
	if i == nil {
		return errors.New("i (*box.Info) cannot be nil")
	}

	off, err := s.Seek(0, io.SeekCurrent)
	if err != nil {
		return err
	}

	i.Offset = uint64(off)
	i.HeaderSize = box.SmallHeader

	buf := bytes.NewBuffer(make([]byte, 0, i.HeaderSize))
	if _, err := io.CopyN(buf, s, int64(i.HeaderSize)); err != nil {
		return err
	}

	data := buf.Bytes()
	i.Size = uint64(binary.BigEndian.Uint32(data))
	i.Type = boxtype.New([4]byte{data[4], data[5], data[6], data[7]})

	if i.Size == 0 {
		off, _ = s.Seek(0, io.SeekEnd)
		i.Size = uint64(off) - i.Offset
		i.ExtendToEOF = true
		if _, err := s.SeekPayload(i); err != nil {
			return err
		}

		return nil
	}

	if i.Size == 1 {
		headerSize := box.LargeHeader - box.SmallHeader
		buf.Reset()
		if _, err := io.CopyN(buf, s, int64(headerSize)); err != nil {
			return err
		}

		i.HeaderSize += headerSize
		i.Size = binary.BigEndian.Uint64(buf.Bytes())

		return nil
	}

	return nil
}

func (s *scanner) SeekPayload(info *box.Info) (int64, error) {
	return s.Seek(int64(info.Offset+info.HeaderSize), io.SeekStart)
}
