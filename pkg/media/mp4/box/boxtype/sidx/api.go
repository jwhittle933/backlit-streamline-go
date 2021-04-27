package sidx

const (
	SIDX string = "sidx"
)

type Box struct {
	ReferenceID                uint32
	Timescale                  uint32
	EarliestPresentationTimeV0 uint32
	FirstOffsetV0              uint32
	EarliestPresentationTimeV1 uint64
	FirstOffsetV1              uint64
	_reserved                  uint16
	ReferenceCount             uint16
	References                 []Reference
}

type Reference struct {
	ReferenceType      bool
	ReferencedSize     uint32
	SubsegmentDuration uint32
	StartsWithSAP      bool
	SAPType            uint32
	SAPDeltaTime       uint32
}

func (Box) Type() string {
	return SIDX
}
