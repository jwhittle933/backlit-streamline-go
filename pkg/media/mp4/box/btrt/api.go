package btrt

const (
	BTRT string = "btrt"
)

type Box struct {
	BufferSizeDB uint32
	MaxBitrate   uint32
	AvgBitrate   uint32
}

func (Box) Type() string {
	return BTRT
}
