package elst

const (
	ELST string = "elst"
)

// Elst is ISOBMFF elst box type
type Elst struct {
	EntryCount uint32
	Entries    []Entry
}

type Entry struct {
	SegmentDurationV0 uint32
	MediaTimeV0       int32
	SegmentDurationV1 uint64
	MediaTimeV1       int64
	MediaRateInteger  int16
	MediaRateFraction int16
}

func (Elst) Type() string {
	return ELST
}
