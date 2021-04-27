package pasp

const (
	PASP string = "pasp"
)

type Box struct {}

func (Box) Type() string {
	return PASP
}