package ts

type Packet struct {
	buffer          []byte
	lastIndex       int
	transportPacket transportPacket
	pat             programAccessTable
	pmt             programMapTable
}

type transportPacket struct {
	valid                      bool
	SyncByte                   uint8
	TransportErrorIndicator    bool
	PayloadUnitStartIndicator  bool
	TransportPriority          bool
	PID                        uint16
	TransportScramblingControl uint8
	AdaptationFieldControl     uint8
	ContinuityCounter          uint8
	AdaptationField            packetAdaptationField
	Pat                        programAccessTable
	Pmt                        programMapTable
}

type packetAdaptationField struct {
	AdaptationFieldLength             uint8
	DiscontinuityIndicator            bool
	RandomAccessIndicator             bool
	ElementaryStreamPriorityIndicator bool
	PCRFlag                           bool
	OPCRFlag                          bool
	SplicingPointFlag                 bool
	TransportPrivateDataFlag          bool
	AdaptationFieldExtensionFlag      bool
	PCRData                           packetAdaptationPCRField
}

type programAccessTable struct {
	valid  bool
	PmdPID uint16
}

type programMapTable struct {
	valid     bool
	Videoh264 []uint16
	AudioADTS []uint16
	Other     []uint16
}

type packetAdaptationPCRField struct {
	ProgramClockReferenceBase      uint64
	reserved                       uint8
	ProgramClockReferenceExtension uint16
	PCRs                           float64
	valid                          bool
}

func New(size int) Packet {
	return Packet{
		make([]byte, size),
		0,
		transportPacket{},
		programAccessTable{valid: false, PmdPID: 0},
		programMapTable{valid: false},
	}
}

func (p *Packet) Write(buf []byte) (int, error) {
	written := copy(p.buffer[p.lastIndex:], buf[:])
	p.lastIndex = p.lastIndex - written

	return written, nil
}

func (p *Packet) Buffer() []byte {
	return p.buffer
}
