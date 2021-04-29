// Package mp4 for MP4 parsing
// See: https://dev.to/sunfishshogi/go-mp4-golang-library-and-cli-tool-for-mp4-52o1
// See: https://openmp4file.com/format.html#:~:text=MP4%20structures%20are%20typically%20referred,below%20have%20precisely%204%20symbols.
// See: https://www.ramugedia.com/mp4-container
package mp4

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/header"
	"github.com/jwhittle933/streamline/pkg/result"
)

type MP4 struct {
	r     io.ReadSeeker
	Size  header.Sizer
	Type  header.Sizer
	Boxes []box.Boxed
}

func New(r io.ReadSeeker) (*MP4, error) {
	res := result.NewSuccess(&MP4{r: r})

	return res.Success.(*MP4), res.Error
}

func withSize(data interface{}) *result.Result {
	mp4 := data.(*MP4)
	return result.NewSuccess(mp4)
}

func (mp4 *MP4) withType() *MP4 {
	return mp4
}

func (mp4 *MP4) Offset() (int64, error) {
	return mp4.r.Seek(0, io.SeekCurrent)
}

// Read satisfies io.Reader interface
func (mp4 *MP4) Read(p []byte) (int, error) {
	return mp4.r.Read(p)
}

// Seek satisfies the io.Seeker interface
func (mp4 *MP4) Seek(offset int64, whence int) (int64, error) {
	return mp4.r.Seek(offset, whence)
}

// JSON encodes mp4 to JSON representation
func (mp4 *MP4) JSON() string {
	return "(*MP4).JSON() unimplemented"
}

func (mp4 *MP4) ReadInfo() (*box.Info, error) {
	off, _ := mp4.Offset()

	bi := &box.Info{
		Offset:     uint64(off),
		HeaderSize: box.SmallHeader,
	}

	buf := bytes.NewBuffer(make([]byte, 0, bi.HeaderSize))
	if _, err := io.CopyN(buf, mp4, box.SmallHeader); err != nil {
		return nil, err
	}

	data := buf.Bytes()
	bi.Size = uint64(binary.BigEndian.Uint32(data))
	bi.Type = [4]byte{data[4], data[5], data[6], data[7]}

	if bi.Size == 0 {
		off, _ = mp4.r.Seek(0, io.SeekEnd)
		bi.Size = uint64(off) - bi.Offset
		bi.ExtendToEOF = true
		if _, err := bi.SeekPayload(mp4); err != nil {
			return nil, err
		}

		return bi, nil
	}

	if bi.Size == 1 {
		buf.Reset()
		if _, err := io.CopyN(buf, mp4, box.LargeHeader-box.SmallHeader); err != nil {
			return nil, err
		}

		bi.HeaderSize += box.LargeHeader - box.SmallHeader
		bi.Size = binary.BigEndian.Uint64(buf.Bytes())
		return bi, nil
	}

	_, err := mp4.Seek(int64(bi.Size), io.SeekStart)
	return bi, err
}
