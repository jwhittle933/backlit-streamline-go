package decoder

const (
	BaselineProfile uint8 = 66
	MainProfile     uint8 = 77
	ExtendedProfile uint8 = 88
	HighProfile     uint8 = 100
	High10Profile   uint8 = 110
	High422Profile  uint8 = 122
)

type Configuration struct {
	Version                    uint8
	Profile                    uint8
	ProfileCompatibility       uint8
	Level                      uint8
	_reserved                  uint8
	LengthSizeMinusOne         uint8
	_reserved2                 uint8
	SequenceParameterSetsLen   uint8
	SequenceParameterSets      []ParameterSet
	PictureParameterSetsLen    uint8
	PictureParameterSets       []ParameterSet
	HighProfileFieldsEnabled   bool
	_reserved3                 uint8
	ChromaFormat               uint8
	_reserved4                 uint8
	BitDepthLumaMinus8         uint8
	_reserved5                 uint8
	BitDepthChromaMinus8       uint8
	SequenceParameterSetExtLen uint8
	SequenceParameterSetExts   []ParameterSet
}

type ParameterSet struct {
	Length  uint16
	NALUnit []byte
}
