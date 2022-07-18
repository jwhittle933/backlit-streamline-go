package sdtp

const (
	SDTP string = "sdtp"
)

type Box struct {
	IsLeading           uint8
	SampleDependsOn     uint8
	SampleIsDependedOn  uint8
	SampleHasRedundancy uint8
}

func (Box) Type() string {
	return SDTP
}
