package slicereader

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader_Uint8(t *testing.T) {
	//
}

func TestReader_Uint16(t *testing.T) {
	//
}

func TestReader_Uint32(t *testing.T) {
	//
}

func TestReader_Uint64(t *testing.T) {
	//
}

func TestReader_Skip(t *testing.T) {
	//
}

func TestReader_Slice(t *testing.T) {
	tests := []struct {
		name     string
		readLen  int
		expected []byte
		setup    func() *Reader
	}{
		{
			name:     "Slice(3) should return the first 3 bytes",
			readLen:  3,
			expected: []byte{1, 255, 31},
			setup: func() *Reader {
				return New([]byte{1, 255, 31, 30, 2, 22})
			},
		},
		{
			name:     "Slice(3) after 1 byte read should return [1:4]",
			readLen:  3,
			expected: []byte{255, 31, 30},
			setup: func() *Reader {
				sr := New([]byte{1, 255, 31, 30, 2, 22})
				sr.Uint8()

				return sr
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sr := test.setup()
			actual := sr.Slice(test.readLen)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestReader_String(t *testing.T) {
	tests := []struct {
		name     string
		readLen  int
		expected string
		setup    func() *Reader
	}{
		{
			name:     "String(9) should return the first 9 characters",
			readLen:  9,
			expected: "comealong",
			setup: func() *Reader {
				return New([]byte("comealongcedric"))
			},
		},
		{
			name:     "Skip(1).String(9) should return the first 9 characters offset 1",
			readLen:  9,
			expected: "omealongc",
			setup: func() *Reader {
				return New([]byte("comealongcedric")).Skip(1)
			},
		},
		{
			name:     "Reader with error should empty string",
			readLen:  9,
			expected: "",
			setup: func() *Reader {
				return &Reader{err: errors.New("this is an error")}
			},
		},
		{
			name:     "Read read past Length should return empty string",
			readLen:  20,
			expected: "",
			setup: func() *Reader {
				return New([]byte("comealongcedric"))
			},
		},
		{
			name:     "AVC Config Compressor Name",
			readLen:  int(uint8(14)),
			expected: "compressorname",
			setup: func() *Reader {
				sr := New(
					append([]byte{14}, []byte("compressorname")...),
				)

				sr.Uint8()
				return sr
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sr := test.setup()
			actual := sr.String(test.readLen)

			assert.Equal(t, test.expected, actual)
		})
	}
}
