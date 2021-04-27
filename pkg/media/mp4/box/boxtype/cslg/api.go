package cslg

const (
	CSLG string = "cslg"
)

type Cslg struct {
	CompositionToDTSShiftV0        int32
	LeastDecodeToDisplayDeltaV0    int32
	GreatestDecodeToDisplayDeltaV0 int32
	CompositionStartTimeV0         int32
	CompositionEndTimeV0           int32
	CompositionToDTSShiftV1        int32
	LeastDecodeToDisplayDeltaV1    int32
	GreatestDecodeToDisplayDeltaV1 int32
	CompositionStartTimeV1         int32
	CompositionEndTimeV1           int32
}

func (c Cslg) Type() string {
	return CSLG
}
