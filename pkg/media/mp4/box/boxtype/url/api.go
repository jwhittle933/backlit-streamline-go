package url

const (
	URL  string = "url "
	SelfContained uint = 0x000001
)

type Url struct {
	Location string
}

func (u Url) Type() string {
	return URL
}
