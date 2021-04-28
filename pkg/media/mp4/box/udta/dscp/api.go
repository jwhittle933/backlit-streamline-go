package dscp

const (
	DSCP string = "dscp"
)

type Box struct {
	//
}

func (Box) Type() string {
	return DSCP
}
