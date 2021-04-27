package btrt

const (
	BTRT string = "btrt"
)

type Btrt struct {
	BufferSizeDB uint32
	MaxBitrate   uint32
	AvgBitrate   uint32
}

func (b Btrt) Type() string {
	return BTRT
}
