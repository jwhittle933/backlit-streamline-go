package envc

const (
	ENCV string = "encv"
)

type Box struct {
	//
}

func (Box) Type() string {
	return ENCV
}
