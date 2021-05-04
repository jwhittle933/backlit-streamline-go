package flag

import (
	"encoding/binary"
	"encoding/hex"
)

type Flag [4]byte

func New(flag [4]byte) Flag {
	return flag
}

func (f Flag) Hex() string {
	return hex.Dump(f[:])
}

func (f Flag) Uint32() uint32 {
	return binary.BigEndian.Uint32(f[:])
}
