package url

const (
	URL  string = "url "
	SelfContained uint = 0x000001
)

type Box struct {
	Location string
}

func (Box) Type() string {
	return URL
}
