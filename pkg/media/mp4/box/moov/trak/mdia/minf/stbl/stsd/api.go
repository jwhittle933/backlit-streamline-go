// Package stsd (Sample Descriptions)
package stsd

const (
	STSD string = "stsd"
)

type Box struct {
	EntryCount uint32
}

func (Box) Type() string {
	return STSD
}
