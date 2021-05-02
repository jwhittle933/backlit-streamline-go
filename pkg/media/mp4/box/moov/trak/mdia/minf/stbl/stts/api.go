// Package stts (Time to Sample) decoding
package stts

const (
	STTS string = "stts"
)

type Box struct {
	EntryCount uint32
	Entries    []Entry
}

type Entry struct {
	SampleCount uint32
	SampleDelta uint32
}

func (Box) Type() string {
	return STTS
}
