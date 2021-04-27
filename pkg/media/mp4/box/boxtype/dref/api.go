package dref

const (
	DREF string = "dref"
)

// Dref is ISOBMFF dref box type
type Dref struct {
	EntryCount uint32
}

func (d Dref) Type() string {
	return DREF
}

