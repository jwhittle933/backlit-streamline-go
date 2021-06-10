package slicereader

import (
	"encoding/binary"
	"errors"
	"fmt"
)

var ErrorSliceRead = errors.New("read too far in SliceReader")

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

	ret := r.slice[r.pos:n]
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

func (r *Reader) checkLen(max int) {
	if r.pos > max {
		r.err = ErrorSliceRead
	}
}
