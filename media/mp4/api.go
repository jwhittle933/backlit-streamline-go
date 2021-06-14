// Package mp4 for MP4 Reading and Writing
package mp4

import (
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/jwhittle933/streamline/media/mp4/box"
	"github.com/jwhittle933/streamline/media/mp4/box/children"
	"github.com/jwhittle933/streamline/media/mp4/box/free"
	"github.com/jwhittle933/streamline/media/mp4/box/ftyp"
	"github.com/jwhittle933/streamline/media/mp4/box/mdat"
	"github.com/jwhittle933/streamline/media/mp4/box/moof"
	"github.com/jwhittle933/streamline/media/mp4/box/moov"
	"github.com/jwhittle933/streamline/media/mp4/box/sidx"
	"github.com/jwhittle933/streamline/media/mp4/box/styp"
	"github.com/jwhittle933/streamline/media/mp4/box/unknown"
	"github.com/jwhittle933/streamline/result"
)

var Children = children.Registry{
	ftyp.FTYP: ftyp.New,
	mdat.MDAT: mdat.New,
	moov.MOOV: moov.New,
	moof.MOOF: moof.New,
	styp.STYP: styp.New,
	free.FREE: free.New,
	sidx.SIDX: sidx.New,
	//emsg.EMSG: emsg.New,
}

type MP4 struct {
	r     io.ReadSeeker
	Size  uint64
	Boxes []box.Boxed
}

func New(r io.ReadSeeker) (*MP4, error) {
	res := result.Wrap(&MP4{r: r, Size: 0, Boxes: make([]box.Boxed, 0)})

	return res.Ok().(*MP4), res.Err()
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
	s := fmt.Sprintf("[\033[1;35mmp4\033[0m] size=%d, boxes=%d\n", m.Size, len(m.Boxes))

	for _, b := range m.Boxes {
		s += b.String()
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
	bi := &box.Info{}
	err := box.ScanInfo(m, bi)
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
// from the mp4's underlying reader, and passes
// reading responsibility for each box's children
// to the Box
func (m *MP4) ReadAll() error {
	var b box.Boxed
	var err error

	for {
		b, err = m.ReadNext()
		if err == io.EOF {
			return nil
		}

		if err != nil || b == nil {
			break
		}

		m.Boxes = append(m.Boxes, b)
	}

	return err
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
