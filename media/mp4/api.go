// Package mp4 for MP4 Reading and Writing
package mp4

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/emsg"
	"github.com/jwhittle933/streamline/media/mp4/box/free"
	"github.com/jwhittle933/streamline/media/mp4/box/ftyp"
	"github.com/jwhittle933/streamline/media/mp4/box/mdat"
	"github.com/jwhittle933/streamline/media/mp4/box/mfra"
	"github.com/jwhittle933/streamline/media/mp4/box/moof"
	"github.com/jwhittle933/streamline/media/mp4/box/moov"
	"github.com/jwhittle933/streamline/media/mp4/box/scanner"
	"github.com/jwhittle933/streamline/media/mp4/box/sidx"
	"github.com/jwhittle933/streamline/media/mp4/box/styp"
	"github.com/jwhittle933/streamline/media/mp4/box/unknown"
	"github.com/jwhittle933/streamline/media/mp4/children"
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
	}
)

var (
	ErrorNotMP4 = errors.New("invalid file type")
)

type MP4 struct {
	r          io.ReadSeeker
	Size       uint64      `json:"size"`
	Emsg       *emsg.Box   `json:"emsg"`
	Free       *free.Box   `json:"free"`
	Ftyp       *ftyp.Box   `json:"ftyp"`
	Mdat       *mdat.Box   `json:"mdat"`
	Mfra       *mfra.Box   `json:"mfra"`
	Moof       *moof.Box   `json:"moof"`
	Moov       *moov.Box   `json:"moov"`
	Sidx       *sidx.Box   `json:"sidx"`
	Styp       *styp.Box   `json:"styp"`
	Children   []box.Boxed `json:"boxes"`
	fragmented bool
}

func New(r io.ReadSeeker) *MP4 {
	return &MP4{r: r, Size: 0, Children: make([]box.Boxed, 0)}
}

func Open(path string) result.Result {
	p, err := filepath.Abs(path)
	if err != nil {
		return result.WrapErr(err)
	}

	if ext := filepath.Ext(p); ext != ".mp4" {
		return result.WrapErr(ErrorNotMP4)
	}

	file, err := os.Open(p)
	if err != nil {
		return result.WrapErr(err)
	}

	return result.Wrap(New(file))
}

func (m *MP4) Offset() (int64, error) {
	return m.r.Seek(0, io.SeekCurrent)
}

// Read satisfies io.Reader interface
func (m *MP4) Read(p []byte) (int, error) {
	return m.r.Read(p)
}

// Seek satisfies the io.Seeker interface
func (m *MP4) Seek(offset int64, whence int) (int64, error) {
	return m.r.Seek(int64(m.Size), whence)
}

func (m *MP4) String() string {
	s := fmt.Sprintf("[\033[1;35mmp4\033[0m] size=%d, fragmented%+v, boxes=%d\n", m.Size, m.fragmented, len(m.Children))

	for _, b := range m.Children {
		s += fmt.Sprintf("%s", b)
	}

	return s
}

// JSON encodes mp4 to JSON representation
func (m *MP4) JSON() string {
	return "(*MP4).JSON() unimplemented"
}

// Hex hex dumps the mp4
func (m *MP4) Hex() string {
	src, _ := ioutil.ReadAll(m)

	return hex.Dump(src)
}

// ReadNext reads and returns the next Box from
// the mp4's underlying reader
func (m *MP4) ReadNext() (box.Boxed, error) {
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

	m.AddBox(b)
	return b, nil
}

// ReadAll iteratively reads every top-level Box
// from the mp4's underlying reader, and passes
// reading responsibility for each box's children
// to the Box
func ReadAll(m *MP4) error {
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
// from the mp4's underlying reader, and passes
// reading responsibility for each box's children
// to the Box
func (m *MP4) ReadAll() error {
	return ReadAll(m)
}

func (m *MP4) AddBox(b box.Boxed) {
	switch b.Info().Type.String() {
	case "emsg":
		m.Emsg = b.(*emsg.Box)
	case "free":
		m.Free = b.(*free.Box)
	case "ftyp":
		m.Ftyp = b.(*ftyp.Box)
	case "mdat":
		m.Mdat = b.(*mdat.Box)
	case "mfra":
		m.Mfra = b.(*mfra.Box)
	case "moof":
		m.Moof = b.(*moof.Box)
	case "moov":
		m.Moov = b.(*moov.Box)
	case "sidx":
		m.Sidx = b.(*sidx.Box)
	case "styp":
		m.Styp = b.(*styp.Box)
	}
}

// Valid returns the validity of the mp4
// based on ISO BMFF standards
// 1. ftyp must contain major brand or compatible brands unsupported by the user agent
// 2. A box or field in the moov box violates the requirements mandated by the major brand
//    or one of the compatible brands
// 3. The tracks in moov contain samples (i.e., the entry count in stts, stsc, or stco boxes
//    are not set to 0
// 4. The Movie Extends box (mvex) is not located in the Movie Box (moov) to indicate that
//    Movie Fragments are to be expected
func (m *MP4) Valid() bool {
	return true
}

func (m *MP4) IsFragmented() bool {
	return m.fragmented
}
