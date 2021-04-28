package colr

const (
	COLR string = "colr"
)

type Box struct {
	ColorType               [4]byte
	ColorPrimaries          uint16
	TransferCharacteristics uint16
	MatrixCoefficients      uint16
	FullRangeFlag           bool
	Reserved                uint8
	Profile                 []byte
	Unknown                 []byte
}

func (Box) Type() string {
	return COLR
}

func (b Box) IsOptFieldEnabled(name string) bool {
	switch b.ColorType {
	case [4]byte{'n', 'c', 'l', 'x'}:
		switch name {
		case "ColorType", "ColorPrimaries", "TransferCharacteristics", "MatrixCoefficients", "FullRangeFlag", "Reserved":
			return true
		default:
			return false
		}
	case [4]byte{'r', 'I', 'C', 'C'}, [4]byte{'p', 'r', 'o', 'f'}:
		return name == "Profile"
	default:
		return name == "Unknown"
	}
}
