package sbgp

const (
	SBGP string = "sbgp"
)

type Box struct {
	GroupingType          uint32
	GroupingTypeParameter uint32
	EntryCount            uint32
	Entries               []Entry
}

type Entry struct {
	SampleCount           uint32
	GroupDescriptionIndex uint32
}

func (Box) Type() string {
	return SBGP
}
