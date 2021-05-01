// Package mehd for ISO BMFF Movie Extends Header
package mehd

const (
	MEHD string = "mehd"
)

type Box struct {
	FragmentDurationV0 uint32
	FragmentDurationV1 uint64
}

func (Box) Type() string {
	return MEHD
}
