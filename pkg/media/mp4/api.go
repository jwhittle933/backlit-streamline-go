// Package mp4 for MP4 parsing
// See: https://dev.to/sunfishshogi/go-mp4-golang-library-and-cli-tool-for-mp4-52o1
// See: https://openmp4file.com/format.html#:~:text=MP4%20structures%20are%20typically%20referred,below%20have%20precisely%204%20symbols.
// See: https://www.ramugedia.com/mp4-container
package mp4

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/header"
	"github.com/jwhittle933/streamline/pkg/result"
)

type MP4 struct {
	raw   []byte
	Size  header.Sizer
	Type  header.Sizer
	Boxes []box.Box
}

func New(src []byte) (*MP4, error) {
	r := result.
		NewSuccess(&MP4{raw: src}).
		Next(withSize)

	return r.Success.(*MP4), r.Error
}

func withSize(data interface{}) *result.Result {
	mp4 := data.(*MP4)
	h, err := header.New(mp4.raw)
	if err != nil {
		result.NewError(err)
	}

	mp4.Size = h
	return result.NewSuccess(mp4)
}

func (mp4 *MP4) withType() *MP4 {
	return mp4
}
