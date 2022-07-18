// Package mpeg for MPEG Reading and Writing
package mpeg

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/box/emsg"
	"github.com/jwhittle933/streamline/media/mpeg/box/free"
	"github.com/jwhittle933/streamline/media/mpeg/box/ftyp"
	"github.com/jwhittle933/streamline/media/mpeg/box/mdat"
	"github.com/jwhittle933/streamline/media/mpeg/box/mfra"
	"github.com/jwhittle933/streamline/media/mpeg/box/moof"
	"github.com/jwhittle933/streamline/media/mpeg/box/moov"
	"github.com/jwhittle933/streamline/media/mpeg/box/moov/pssh"
	"github.com/jwhittle933/streamline/media/mpeg/box/scanner"
	"github.com/jwhittle933/streamline/media/mpeg/box/sidx"
	"github.com/jwhittle933/streamline/media/mpeg/box/styp"
	"github.com/jwhittle933/streamline/media/mpeg/box/unknown"
	"github.com/jwhittle933/streamline/media/mpeg/children"
	"github.com/jwhittle933/streamline/result"
)

var (
	Children = children.Registry{
		emsg.EMSG: emsg.New,
		free.FREE: free.New,
		ftyp.FTYP: ftyp.New,
		mdat.MDAT: mdat.New,
		mfra.MFRA: mfra.New,
		moof.MOOF: moof.New,
		moov.MOOV: moov.New,
		sidx.SIDX: sidx.New,
		styp.STYP: styp.New,
		pssh.PSSH: pssh.New,
	}
)

var (
	ErrorNotMP4 = errors.New("invalid file type")
)

type MPEG struct {
	r          io.ReadSeeker
	Size       uint64      `json:"size"`
	Children   []box.Boxed `json:"boxes"`
	fragmented bool
}

func New(r io.ReadSeeker) *MPEG {
	return &MPEG{
		r:        r,
		Size:     0,
		Children: make([]box.Boxed, 0),
	}
}

func Open(path string) result.Result {
	p, err := filepath.Abs(path)
	if err != nil {
		return result.WrapErr(err)
	}

	file, err := os.Open(p)
	if err != nil {
		return result.WrapErr(err)
	}

	return result.Wrap(New(file))
}

func (m *MPEG) Offset() (int64, error) {
	return m.r.Seek(0, io.SeekCurrent)
}

// Read satisfies io.Reader interface
func (m *MPEG) Read(p []byte) (int, error) {
	return m.r.Read(p)
}

// Seek satisfies the io.Seeker interface
func (m *MPEG) Seek(offset int64, whence int) (int64, error) {
	return m.r.Seek(int64(m.Size), whence)
}

func (m *MPEG) String() string {
	s := fmt.Sprintf("[\033[1;35mMP4\033[0m] size=%d, fragmented=%+v, boxes=%d\n",
		m.Size,
		m.fragmented,
		len(m.Children),
	)

	for _, b := range m.Children {
		s += fmt.Sprintf("%s\n", b)
	}

	return s
}

// JSON encodes mpeg to JSON representation
func (m *MPEG) JSON() string {
	return "(*MPEG).JSON() unimplemented"
}

// Hex dumps the mpeg in hex format
func (m *MPEG) Hex() string {
	src, _ := ioutil.ReadAll(m)

	return hex.Dump(src)
}

// ReadNext reads and returns the next Box from
// the mpeg's underlying reader
func (m *MPEG) ReadNext() (box.Boxed, error) {
	sc := scanner.New(m)

	bi := &box.Info{}
	err := sc.ScanInfo(bi)
	if err != nil {
		return nil, err
	}

	var boxFactory box.Factory
	var found bool
	if boxFactory, found = Children[bi.Type.String()]; !found {
		boxFactory = unknown.New
	}

	b := boxFactory(bi)

	// copy bytes starting at offset until Size - HeaderSize
	if _, err := io.CopyN(b, m, int64(bi.Size-bi.HeaderSize)); err != nil {
		return nil, err
	}

	return b, nil
}

// ReadAll iteratively reads every top-level Box
// from the mpeg's underlying reader, and passes
// reading responsibility for each box's children
// to the Box
func ReadAll(m *MPEG) error {
	var b box.Boxed
	var err error

	for {
		b, err = m.ReadNext()
		if err == io.EOF {
			for _, c := range m.Children {
				m.Size += c.Info().Size
			}

			return nil
		}

		if err != nil || b == nil {
			break
		}

		switch b.Info().Type.String() {
		case "moof", "styp":
			m.fragmented = true
		}

		m.Children = append(m.Children, b)
	}

	return err
}

// ReadAll iteratively reads every top-level Box
// from the mpeg's underlying reader, and passes
// reading responsibility for each box's children
// to the Box
func (m *MPEG) ReadAll() error {
	return ReadAll(m)
}

// Valid returns the validity of the mpeg
// based on ISO BMFF standards
// 1. ftyp must contain major brand or compatible brands unsupported by the user agent
// 2. A box or field in the moov box violates the requirements mandated by the major brand
//    or one of the compatible brands
// 3. The tracks in moov contain samples (i.e., the entry count in stts, stsc, or stco boxes
//    are not set to 0
// 4. The Movie Extends box (mvex) is not located in the Movie Box (moov) to indicate that
//    Movie Fragments are to be expected
func (m *MPEG) Valid() bool {
	return true
}

func (m *MPEG) IsFragmented() bool {
	return m.fragmented
}
