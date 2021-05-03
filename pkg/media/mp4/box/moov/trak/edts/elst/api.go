// Pacakge elst (Edit List Box)
package elst

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/base"
)

const (
	ELST string = "elst"
)

// Box is ISOBMFF elst box type
type Box struct {
	base.Box
	EntryCount uint32
	Entries    []Entry
}

type Entry struct {
	SegmentDurationV0 uint32
	MediaTimeV0       int32
	SegmentDurationV1 uint64
	MediaTimeV1       int64
	MediaRateInteger  int16
	MediaRateFraction int16
}

func New(i *box.Info) box.Boxed {
	return &Box{base.Box{BoxInfo: i}, 0, make([]Entry, 0)}
}

func (Box) Type() string {
	return ELST
}

func (b Box) String() string {
	return b.Info().String()
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	return len(src), nil
}
