package udta

const (
	UDTA string = "udta"
)

// Box is ISOBMFF udta box type
type Box struct {
	//
}

type Udat3GPPString struct {
	Pad      bool
	Language [3]byte
	Data     []byte
}

func (Box) Type() string {
	return UDTA
}
