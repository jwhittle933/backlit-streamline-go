package perf

const (
	PERF string = "perf"
)

type Box struct {
	//
}

func (Box) Type() string {
	return PERF
}
