package box

import "io"

type Boxed interface {
	io.Writer
	Type() string
	Version() uint8
	Children() []Boxed
}

type Info struct {
	Offset      uint64
	Size        uint64
	HeaderSize  uint64
	ExtendToEOF bool
}
