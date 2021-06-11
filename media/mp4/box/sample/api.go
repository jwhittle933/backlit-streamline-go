package sample

import (
	"github.com/jwhittle933/streamline/media/mp4/box/base"
)

type Entry struct {
	base.Box
	_reserved          [6]uint8
	DataReferenceIndex uint16 // 2 bytes
}
