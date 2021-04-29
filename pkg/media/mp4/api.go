// Package mp4 for MP4 parsing
// See: https://dev.to/sunfishshogi/go-mp4-golang-library-and-cli-tool-for-mp4-52o1
// See: https://openmp4file.com/format.html#:~:text=MP4%20structures%20are%20typically%20referred,below%20have%20precisely%204%20symbols.
// See: https://www.ramugedia.com/mp4-container
package mp4

import (
	"io"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/header"
	"github.com/jwhittle933/streamline/pkg/result"
)

type MP4 struct {
	raw   []byte
	Size  header.Sizer
	Type  header.Sizer
	Boxes []box.Boxed
}

func New(r io.ReadSeeker) (*MP4, error) {
	buf := make([]byte, 4)
	_, _ = r.Read(buf)

	res := result.NewSuccess(&MP4{raw: buf})

	return res.Success.(*MP4), res.Error
}

func withSize(data interface{}) *result.Result {
	mp4 := data.(*MP4)
	return result.NewSuccess(mp4)
}

func (mp4 *MP4) withType() *MP4 {
	return mp4
}
