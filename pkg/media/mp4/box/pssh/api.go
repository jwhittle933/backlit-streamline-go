package pssh

const (
	PSSH string = "pssh"
)

// Box is ISOBMFF pssh box type
type Box struct {
	SystemID [16]byte
	KIDCount uint32
	KIDs     []KID
	DataSize int32
	Data     []byte
}

type KID [16]byte

func (Box) Type() string {
	return PSSH
}
