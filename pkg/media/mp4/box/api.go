package box

type Box interface {
	Type() string
	Version() uint8
}

type Info struct {
	Offset      uint64
	Size        uint64
	HeaderSize  uint64
	ExtendToEOF bool
}
