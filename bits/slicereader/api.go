package slicereader

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

var (
	ErrorSliceRead  = errors.New("read too far in SliceReader")
	ErrorSeekTooFar = errors.New("seek requested beyond length of reader")
)

type Reader struct {
	slice []byte
	pos   int
	len   int
	err   error
}

func New(s []byte) *Reader {
	return &Reader{
		slice: s,
		pos:   0,
		len:   len(s),
		err:   nil,
	}
}

func (r *Reader) Uint8() byte {
	if r.err != nil {
		return 0
	}

	if r.checkLen(r.len - 1); r.err != nil {
		return 0
	}

	res := r.slice[r.pos]
	r.pos++
	return res
}

func (r *Reader) Uint16() uint16 {
	if r.err != nil {
		fmt.Println("Reader Error")
		return 0
	}

	if r.checkLen(r.len - 2); r.err != nil {
		return 0
	}

	res := binary.BigEndian.Uint16(r.slice[r.pos : r.pos+2])
	r.pos += 2
	return res
}

func (r *Reader) Uint32() uint32 {
	if r.err != nil {
		return 0
	}

	if r.checkLen(r.len - 4); r.err != nil {
		return 0
	}

	res := binary.BigEndian.Uint32(r.slice[r.pos : r.pos+4])
	r.pos += 4
	return res
}

func (r *Reader) Uint64() uint64 {
	if r.err != nil {
		return 0
	}

	if r.checkLen(r.len - 8); r.err != nil {
		return 0
	}

	res := binary.BigEndian.Uint64(r.slice[r.pos : r.pos+8])
	r.pos += 8
	return res
}

func (r *Reader) Int64() int64 {
	return int64(r.Uint64())
}

func (r *Reader) Skip(n int) *Reader {
	if r.err != nil {
		return r
	}

	if r.pos+n > r.Length() {
		r.err = fmt.Errorf("attempt to skip bytes to pos %d beyond slice len %d", r.pos+n, r.len)
		return r
	}

	r.pos += n
	return r
}

func (r *Reader) Slice(n int) []byte {
	if r.err != nil {
		return []byte{}
	}

	if r.pos+n > r.Length() {
		return []byte{}
	}

	ret := r.slice[r.pos : r.pos+n]
	r.pos += n

	return ret
}

func (r *Reader) Copy(buf []byte) int {
	n := copy(r.slice, buf)
	r.pos += n

	return n
}

func (r *Reader) Remaining() []byte {
	if r.err != nil {
		return []byte{}
	}
	res := r.slice[r.pos:]
	r.pos = r.Length()
	return res
}

// NewFromRemaining returns a new Reader from current position
func (r *Reader) NewFromRemaining() *Reader {
	return &Reader{
		slice: r.slice[r.pos:],
		pos:   0,
		err:   nil,
	}
}

func (r *Reader) String(size int) string {
	if r.err != nil {
		return ""
	}

	if r.pos+size > r.Length() {
		r.err = ErrorSliceRead
		return ""
	}

	res := r.slice[r.pos : r.pos+size]
	r.pos += size
	return string(res)
}

func (r *Reader) Length() int {
	return r.len
}

func (r *Reader) Position() int {
	return r.pos
}

func (r *Reader) Consumed() int {
	return len(r.slice[0:r.pos])
}

func (r *Reader) Error() error {
	return r.err
}

func (r *Reader) EOF() bool {
	return r.pos == len(r.slice)-1
}

// Seek satisfies the io.Seeker interface
func (r *Reader) Seek(offset int64, whence int) (int64, error) {
	if offset < 0 {
		return 0, errors.New("negative offset")
	}

	switch whence {
	case io.SeekStart:
		if int(offset) > r.Length() {
			return 0, ErrorSeekTooFar
		}

		r.pos = int(offset)
		return offset, nil
	case io.SeekCurrent:
		if int(offset)+r.pos > r.Length() {
			return 0, ErrorSeekTooFar
		}

		r.pos += int(offset)
		return int64(r.pos), nil
	case io.SeekEnd:
		if int(offset) > r.Length() {
			return 0, ErrorSeekTooFar
		}

		r.pos = r.Length() - int(offset)
		return int64(r.pos), nil
	default:
		return 0, fmt.Errorf("unknown seek start: %d", whence)
	}
}

// Read satisfies the io.Reader interface
func (r *Reader) Read(p []byte) (int, error) {
	copy(p, r.Slice(len(p)))

	if r.Error() != nil {
		return 0, r.Error()
	}

	return len(p), nil
}

// ReadAt satisfies the io.ReaderAt interface
// ReadAt will adjust the reader's position to `off`,
// read `len(p)` bytes from the underlying reader,
// and reapply the readers original position
func (r *Reader) ReadAt(p []byte, off int64) (int, error) {
	originalPosition := r.pos

	if _, err := r.Seek(off, io.SeekStart); err != nil {
		return 0, err
	}

	if  _, err := r.Read(p); err != nil {
		return 0, err
	}

	r.pos = originalPosition
	return len(p), nil
}

func (r *Reader) checkLen(max int) {
	if r.pos > max {
		r.err = ErrorSliceRead
	}
}
