// Package hdlr (Handler) for common file metadata
package hdlr

const (
	HDLR string = "hdlr"
)

// Box is ISOBMFF hdlr box type
type Box struct {
	// PreDefined: component_type of QuickTime
	// pre_defined of ISO-14496 always has 0
	// component_type has mhlr or dhlr
	PreDefined  uint32
	HandlerType [4]byte
	Reserved    [3]uint64
	Name        string
}

func (Box) Type() string {
	return HDLR
}
