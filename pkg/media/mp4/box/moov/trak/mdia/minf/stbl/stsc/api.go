// Package stsc (Sample-to-Chunk)
package stsc

const (
	STSC string = "stsc"
)

type Box struct {
	EntryCount uint32
	Entries    []Entry
}

type Entry struct {
	FirstChunk             uint32
	SamplesPerChunk        uint32
	SampleDescriptionIndex uint32
}

func (Box) Type() string {
	return STSC
}
