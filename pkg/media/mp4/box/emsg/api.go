package emsg

import (
	"io"

	"github.com/jwhittle933/streamline/pkg/bits"

	"github.com/icza/bitio"
)

const (
	EMSG string = "emsg"
)

type Box struct {
	SchemeIdUri           string
	Value                 string
	Timescale             uint32
	PresentationTimeDelta uint32
	PresentationTime      uint64
	EventDuration         uint32
	Id                    uint32
	MessageData           []byte
}

func (Box) Type() string {
	return EMSG
}

func (b *Box) OnReadField(name string, r bitio.Reader) (uint64, bool, error) {
	// check version?
	var err error

	switch name {
	case "SchemeIdUri", "Value":
		return 0, true, nil
	case "MessageData":
		if b.SchemeIdUri, err = bits.ReadString(r); err != nil {
			return 0, false, err
		}

		if b.Value, err = bits.ReadString(r); err != nil {
			return 0, false, err
		}

		return uint64(len(b.SchemeIdUri)+len(b.Value)+2) * 8, false, nil
	default:
		return 0, false, nil
	}
}

func (b *Box) OnWriteField(name string, w io.Writer) (uint64, bool, error) {
	// check version?
	switch name {
	case "SchemeIdUrl", "Value":
		return 0, true, nil
	case "MessageData":
		if err := bits.WriteString(w, b.SchemeIdUri); err != nil {
			return 0, false, err
		}

		if err := bits.WriteString(w, b.Value); err != nil {
			return 0, false, err
		}

		return uint64(len(b.SchemeIdUri)+len(b.Value)+2) * 8, false, nil
	}

	return 0, false, nil
}
