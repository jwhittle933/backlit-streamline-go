package ctts

const (
	CTTS string = "ctts"
)

type Box struct {
	EntryCount uint32
	Entries    []Entry
}

type Entry struct {
	SampleCount    uint32
	SampleOffsetV0 uint32
	SampleOffsetV1 int32
}

func (Box) Type() string {
	return CTTS
}
