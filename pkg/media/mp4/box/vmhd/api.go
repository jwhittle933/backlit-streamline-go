package vmhd

const (
	VMHD string = "vmhd"
)

type Box struct {
	GraphicsMode uint16
	OpColor      [3]uint16
}

func (Box) Type() string {
	return VMHD
}
