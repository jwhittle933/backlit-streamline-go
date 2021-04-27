package es

type Descriptor struct {
	ESID                 uint16
	StreamDependenceFlag bool
	UrlFlag              bool
	OcrStreamFlag        bool
	StreamPriority       int8
	DependsOnESID        uint16
	URLLength            uint8
	URLString            []byte
	OCRESID              uint16
}

func (d *Descriptor) IsOptFieldEnabled(name string) bool {
	switch name {
	case "DependsOnESID":
		return d.StreamDependenceFlag
	case "URLLength", "URLString":
		return d.UrlFlag
	case "OCRESID":
		return d.OcrStreamFlag
	}

	return false
}
