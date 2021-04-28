package box

import "io"

type Box interface {
	io.Writer
	Type() string
	Version() uint8
}

type Info struct {
	Offset      uint64
	Size        uint64
	HeaderSize  uint64
	ExtendToEOF bool
}
