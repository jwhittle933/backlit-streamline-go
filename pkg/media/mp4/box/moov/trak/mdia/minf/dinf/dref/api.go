// Package dref (Data Reference), declare source of media data in track
package dref

const (
	DREF string = "dref"
)

// Box is ISOBMFF dref box type
type Box struct {
	EntryCount uint32
}

func (Box) Type() string {
	return DREF
}
