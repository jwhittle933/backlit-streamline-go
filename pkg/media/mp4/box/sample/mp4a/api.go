package mp4a

const (
	MP4A string = "mp4a"
)

type Box struct {
	//
}

func (Box) Type() string {
	return MP4A
}
