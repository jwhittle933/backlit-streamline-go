package slicewriter

import (
	"encoding/binary"
	"io"
)

type Writer struct {
	writer io.Writer
}

func New(w io.Writer) *Writer {
	return &Writer{w}
}

func (w *Writer) Uint8(b byte) {
	w.capture(binary.Write(w.writer, binary.BigEndian, b))
}

func (w *Writer) Uint16(b uint16) {
	w.capture(binary.Write(w.writer, binary.BigEndian, b))
}

func (w *Writer) Uint32(b uint32) {
	w.capture(binary.Write(w.writer, binary.BigEndian, b))
}

func (w *Writer) Uint48(b uint64) {
	msb := uint16(b >> 16)
	lsb := uint32(b & 0xffffffff)

	w.capture(binary.Write(w.writer, binary.BigEndian, msb))
	w.capture(binary.Write(w.writer, binary.BigEndian, lsb))
}

func (w *Writer) Uint64(b uint64) {
	w.capture(binary.Write(w.writer, binary.BigEndian, b))
}

func (w *Writer) Zero(delta int) {
	var data byte = 0

	for i := 0; i < delta; i++ {
		w.capture(binary.Write(w.writer, binary.BigEndian, data))
	}
}

func (w *Writer) Slice(b []byte) {
	_, err := w.writer.Write(b)
	w.capture(err)
}

func (w *Writer) capture(err error) {
	if err != nil {
		//
	}
}
