package btrt

import (
	"bytes"
	"fmt"
	slicereader2 "github.com/jwhittle933/streamline/bits/slicereader"
	slicewriter2 "github.com/jwhittle933/streamline/bits/slicewriter"
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	base2 "github.com/jwhittle933/streamline/media/mp4/box/base"
)

const (
	BTRT string = "btrt"
)

type Box struct {
	base2.Box
	BufferSizeDB uint32
	MaxBitrate   uint32
	AvgBitrate   uint32
}

func New(i *box2.Info) box2.Boxed {
	return &Box{base2.Box{BoxInfo: i}, 0, 0, 0}
}

func (Box) Type() string {
	return BTRT
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s, buffersizedb=%d, maxbitrate=%d, avgbitrate=%d",
		b.Info(),
		b.BufferSizeDB,
		b.MaxBitrate,
		b.AvgBitrate,
	)
}

func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader2.New(src)
	b.BufferSizeDB = sr.Uint32()
	b.MaxBitrate = sr.Uint32()
	b.AvgBitrate = sr.Uint32()

	return sr.Length(), nil
}

func (b *Box) Read(dst []byte) (int, error) {
	buf := bytes.NewBuffer(dst)
	sw := slicewriter2.New(buf)

	header := make([]byte, b.Info().HeaderSize+b.Info().Size)
	_, _ = b.Info().Read(header)
	sw.Slice(header)
	sw.Uint32(b.BufferSizeDB)
	sw.Uint32(b.MaxBitrate)
	sw.Uint32(b.AvgBitrate)

	return int(b.Info().Size), nil
}
