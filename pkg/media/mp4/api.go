// Package mp4 for MP4 parsing
// See: https://dev.to/sunfishshogi/go-mp4-golang-library-and-cli-tool-for-mp4-52o1
package mp4

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	header2 "github.com/jwhittle933/streamline/pkg/media/mp4/box/header"
	"github.com/jwhittle933/streamline/pkg/result"
)

type MP4 struct {
	raw   []byte
	Size  header2.Sizer
	Type  header2.Sizer
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
	h, err := header2.New(mp4.raw)
	if err != nil {
		result.NewError(err)
	}

	mp4.Size = h
	return result.NewSuccess(mp4)
}

func (mp4 *MP4) withType() *MP4 {
	// boxtype starts with "ftype"
	return mp4
}