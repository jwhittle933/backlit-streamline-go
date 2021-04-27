package decoder

type ConfigDescriptor struct {
	ObjectTypeIndication byte
	StreamType           int8
	UpStream             bool
	Reserved             bool
	BufferSizeDB         uint32
	MaxBitrate           uint32
	AvgBitrate           uint32
}
