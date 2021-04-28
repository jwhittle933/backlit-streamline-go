package schi

const (
	SCHI string = "schi"
)

type Box struct {
	SchemeType    [4]byte
	SchemeVersion uint32
	SchemeUri     []byte
}

func (Box) Type() string {
	return SCHI
}
