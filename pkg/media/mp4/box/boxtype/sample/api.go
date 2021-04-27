package sample

type Entry struct {
	_reserved          [6]uint8
	DataReferenceIndex uint16
}

type VisualEntry struct {
	Entry
	Predefined           uint16
	_reserved            uint16
	Predefined2          [3]uint32
	Width                uint16
	Height               uint16
	HorizontalResolution uint32
	VerticalResolution   uint32
	_reserved2           uint32
	FrameCount           uint16
	CompressorName       [32]byte
	Depth                uint16
	Predefined3          int16
}

type AudioEntry struct {
	Version       uint16
	_reserved     [3]uint16
	ChannelCount  uint16
	SampleSize    uint16
	Predefined    uint16
	_reserved2    uint16
	SampleRate    uint32
	QuickTimeData []byte
}

type PixelAspectRatio struct {
	HSpacing uint32
	VSpacing uint32
}
