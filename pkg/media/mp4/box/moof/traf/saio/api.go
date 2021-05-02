// Package saio (Sample Auxiliary Information Offsets)
package saio

const (
	SAIO string = "saio"
)

type Box struct {
	AuxInfoType          [4]byte
	AuxInfoTypeParameter uint32
	EntryCount           uint32
	OffsetV0             []uint32
	OffsetV1             []uint64
}

func (Box) Type() string {
	return SAIO
}
