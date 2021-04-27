package ctts

const (
	CTTS string = "ctts"
)

type Ctts struct {
	EntryCount uint32
	Entries    []Entry
}

type Entry struct {
	SampleCount    uint32
	SampleOffsetV0 uint32
	SampleOffsetV1 int32
}

func (c Ctts) Type() string {
	return CTTS
}
