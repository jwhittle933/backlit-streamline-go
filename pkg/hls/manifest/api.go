package manifest

type Manifest interface {
	String() string
}

type VOD struct {}
type LiveEvent struct {}
type LiveWindow struct {}

func (v VOD) String() string {
	return "VOD"
}

func (v LiveEvent) String() string {
	return "EVENT"
}

func (v LiveWindow) String() string {
	return "WINDOW"
}
