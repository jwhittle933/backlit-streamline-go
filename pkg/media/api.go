// Package media parses and validates media formats
package media

import "io"

type Parser interface {
	io.Seeker
	io.Reader
	io.Closer
	MediaType() string
}

// Open opens a media file for parsing
func Open(path string) {
	// do the thing
}

