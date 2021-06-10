package emsg

const (
	EMSG string = "emsg"
)

type Box struct {
	SchemeIdUri           string
	Value                 string
	Timescale             uint32
	PresentationTimeDelta uint32
	PresentationTime      uint64
	EventDuration         uint32
	Id                    uint32
	MessageData           []byte
}

func (Box) Type() string {
	return EMSG
}
