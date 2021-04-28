package auth

const (
	AUTH string = "auth"
)

type Box struct {
	//
}

func (Box) Type() string {
	return AUTH
}
