package stss

const (
	STSS string = "stss"
)

type Box struct {
	EntryCount   uint32
	SampleNumber []uint32
}

func (Box) Type() string {
	return STSS
}
