package co64

const (
	CO64 string = "co64"
)

type Co64 struct {
	EntryCount  uint32
	ChunkOffset []uint64
}

func (c Co64) Type() string {
	return CO64
}
