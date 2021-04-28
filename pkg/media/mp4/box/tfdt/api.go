package tfdt

const (
	TFDT string = "tfdt"
)

type Box struct {
	BaseMediaDecodeTimeV0 uint32
	BaseMediaDecodeTimeV1 uint64
}

func (Box) Type() string {
	return TFDT
}
