package stco

const (
	STCO string = "stco"
)

// Box (stco) stores meta-data about video frames
type Box struct {
	EntryCount  uint32
	ChunkOffset []uint32
}

func (Box) Type() string {
	return STCO
}
