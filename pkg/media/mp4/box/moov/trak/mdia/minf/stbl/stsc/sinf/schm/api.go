// Package schm (Scheme Type)
package schm

const (
	SCHM string = "schm"
)

type Box struct {
	SchemeType    [4]byte
	SchemeVersion uint32
	SchemeUri     []byte
}

func (Box) Type() string {
	return SCHM
}
