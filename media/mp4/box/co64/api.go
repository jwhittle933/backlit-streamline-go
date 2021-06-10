package co64

const (
	CO64 string = "co64"
)

type Box struct {
	EntryCount  uint32
	ChunkOffset []uint64
}

func (Box) Type() string {
	return CO64
}
