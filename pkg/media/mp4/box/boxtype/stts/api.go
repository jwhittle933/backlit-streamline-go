package stts

const (
	STTS string = "stts"
)

type Box struct {

}

func (Box) Type() string {
	return STTS
}