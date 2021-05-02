// Package mp4 for MP4 parsing
// See: https://dev.to/sunfishshogi/go-mp4-golang-library-and-cli-tool-for-mp4-52o1
// See: https://openmp4file.com/format.html#:~:text=MP4%20structures%20are%20typically%20referred,below%20have%20precisely%204%20symbols.
// See: https://www.ramugedia.com/mp4-container
// See: https://bitmovin.com/fun-with-container-formats-2/
// See https://www.w3.org/2013/12/byte-stream-format-registry/isobmff-byte-stream-format.html
package mp4

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"io"
	"io/ioutil"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/boxtype"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/children"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/free"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/ftyp"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/mdat"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moof"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/moov"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/styp"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/unknown"
	"github.com/jwhittle933/streamline/pkg/result"
)

var mp4Children = children.Registry{
	ftyp.FTYP: ftyp.New,
	mdat.MDAT: mdat.New,
	moov.MOOV: moov.New,
	moof.MOOF: moof.New,
	styp.STYP: styp.New,
	free.FREE: free.New,
	//sidx.SIDX: sidx.New,
	//emsg.EMSG: emsg.New,
}

type MP4 struct {
	r     io.ReadSeeker
	Size  uint64
	Boxes []box.Boxed
}

func New(r io.ReadSeeker) (*MP4, error) {
	res := result.NewSuccess(&MP4{r: r, Size: 0, Boxes: make([]box.Boxed, 0)})

	return res.Success.(*MP4), res.Error
}

func (mp4 *MP4) Offset() (int64, error) {
	return mp4.r.Seek(0, io.SeekCurrent)
}

// Read satisfies io.Reader interface
func (mp4 *MP4) Read(p []byte) (int, error) {
	return mp4.r.Read(p)
}

// Seek satisfies the io.Seeker interface
func (mp4 *MP4) Seek(offset int64, whence int) (int64, error) {
	mp4.Size += uint64(offset)
	return mp4.r.Seek(int64(mp4.Size), whence)
}

// JSON encodes mp4 to JSON representation
func (mp4 *MP4) JSON() string {
	return "(*MP4).JSON() unimplemented"
}

// Hex hex dumps the mp4
func (mp4 *MP4) Hex() string {
	src, _ := ioutil.ReadAll(mp4)

	return hex.Dump(src)
}

// ReadNext reads and returns the next Box from
// the mp4's underlying reader
func (mp4 *MP4) ReadNext() (box.Boxed, error) {
	off, _ := mp4.Offset()

	bi := &box.Info{
		Offset:     uint64(off),
		HeaderSize: box.SmallHeader,
	}

	buf := bytes.NewBuffer(make([]byte, 0, bi.HeaderSize))
	if _, err := io.CopyN(buf, mp4, int64(bi.HeaderSize)); err != nil {
		return nil, err
	}

	data := buf.Bytes()
	bi.Size = uint64(binary.BigEndian.Uint32(data))
	bi.Type = boxtype.New([4]byte{data[4], data[5], data[6], data[7]})

	var boxFactory children.BoxFactory
	var found bool
	if boxFactory, found = mp4Children[bi.Type.String()]; !found {
		boxFactory = unknown.New
	}

	b := boxFactory(bi)

	//if bi.Size == 0 {
	//	off, _ = mp4.Seek(0, io.SeekEnd)
	//	bi.Size = uint64(off) - bi.Offset
	//	bi.ExtendToEOF = true
	//	if _, err := box.SeekPayload(mp4, b); err != nil {
	//		return nil, err
	//	}
	//
	//	return b, nil
	//}
	//
	//if bi.Size == 1 {
	//	buf.Reset()
	//	if _, err := io.CopyN(buf, mp4, box.LargeHeader-box.SmallHeader); err != nil {
	//		return nil, err
	//	}
	//
	//	bi.HeaderSize += box.LargeHeader - box.SmallHeader
	//	bi.Size = binary.BigEndian.Uint32(buf.Bytes())
	//
	//	return b, nil
	//}

	if _, err := io.CopyN(b, mp4, int64(bi.Size-bi.HeaderSize)); err != nil {
		return nil, err
	}

	_, err := mp4.Seek(int64(bi.Size), io.SeekStart)
	return b, err
}

// ReadAll iteratively reads every top-level Box
// from the mp4's underlying reader, and passes
// reading responsibility for each box's children
// to the Box
func (mp4 *MP4) ReadAll() error {
	var b box.Boxed
	var err error

	for {
		b, err = mp4.ReadNext()
		if err == io.EOF {
			return nil
		}

		if err != nil || b == nil {
			break
		}

		mp4.Boxes = append(mp4.Boxes, b)
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
func (mp4 *MP4) Valid() bool {
	return true
}
