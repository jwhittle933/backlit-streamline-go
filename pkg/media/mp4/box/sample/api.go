package sample

import "github.com/jwhittle933/streamline/pkg/media/mp4/box/base"

type Entry struct {
	base.Box
	_reserved          [6]uint8
	DataReferenceIndex uint16 // 2 bytes
}

type Visual struct { // 80 bytes
	Entry
	Predefined           uint16    // 2 bytes
	_reserved            uint16    // 2 bytes
	Predefined2          [3]uint32 // 12 bytes
	Width                uint16    // 2 bytes
	Height               uint16    // 2 bytes
	HorizontalResolution uint32    // 2 bytes
	VerticalResolution   uint32    // 4 bytes
	_reserved2           uint32    // 4 bytes
	FrameCount           uint16    // 4 bytes
	CompressorName       [32]byte  // 32 bytes
	Depth                uint16    // 2 bytes
	Predefined3          int16     // 2 bytes
}

type Audio struct {
	Entry
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
	base.Box
	HSpacing uint32
	VSpacing uint32
}
