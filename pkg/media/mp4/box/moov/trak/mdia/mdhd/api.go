// Package mdhd (Media Header)
package mdhd

const (
	MDHD string = "mdhd"
)

// Box is ISOBMFF mdhd box type
type Box struct {
	CreationTimeV0     uint32
	ModificationTimeV0 uint32
	CreationTimeV1     uint64
	ModificationTimeV1 uint64
	Timescale          uint32
	DurationV0         uint32
	DurationV1         uint64
	Pad                bool
	Language           [3]byte
	PreDefined         uint16
}

func (Box) Type() string {
	return MDHD
}
