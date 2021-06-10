// Package tfra (Track Fragment Random Access)
package tfra

const (
	TFRA string = "tfra"
)

type Box struct {
	TrackID               uint32
	_reserved             uint32
	LengthSizeOfTrafNum   byte
	LengthSizeOfTrunNum   byte
	LengthSizeOfSampleNum byte
	NumberOfEntries       uint32
	Entries               []Entry
}

type Entry struct {
	TimeV0       uint32
	MoofOffsetV0 uint32
	TimeV1       uint64
	MoofOffsetV1 uint64
	TrafNumber   uint32
	TrunNumber   uint32
	SampleNumber uint32
}

func (Box) Type() string {
	return TFRA
}
