package stco

const (
	STCO string = "stco"
)

type Box struct {
	EntryCount  uint32
	ChunkOffset []uint32
}

func (Box) Type() string {
	return STCO
}
