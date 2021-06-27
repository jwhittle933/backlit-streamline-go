// Package media parses and validates media formats
package media

import (
	"io"
)

type ChildReader interface {
	ReadChildren() []Box
}

type ChildWriter interface {
	WriteChildren([]byte) (int, error)
}

type Box interface {
	io.Reader
	io.Writer
	ChildReader
	ChildWriter
}

type Reader interface {
	ReadAll() error
	ReadNext() (Box, error)
}
