package uuid

import (
	"encoding/hex"
	"fmt"
)

const (
	PlayReady = "9a04f079-9840-4286-ab92-e65be0885f95"
	Widevine  = "edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"
	FairPlay  = "94CE86FB-07FF-4F43-ADB8-93D2FA968CA2"
	VCAS      = "9a27dd82-fde2-4725-8cbc-4234aa06ec09"
	Invalid   = "494e56414c4944"
)

// UUID 16 byte KeyID or SystemID
type UUID [16]byte

func FromBytes(in []byte) UUID {
	var out UUID
	if len(in) != 16 {
		copy(out[:], append([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0}, "INVALID"...))
		return out
	}

	copy(out[:], in)
	return out
}

func (u UUID) FromBytes(in []byte) UUID {
	return FromBytes(in)
}

func (u UUID) String() string {
	h := hex.EncodeToString(u[:])
	return fmt.Sprintf("%s-%s-%s-%s-%s", h[0:8], h[8:12], h[12:16], h[16:20], h[20:32])
}

func (u UUID) SystemName(systemID UUID) string {
	switch systemID.String() {
	case PlayReady:
		return "PlayReady"
	case Widevine:
		return "Widevine"
	case FairPlay:
		return "FairPlay"
	case VCAS:
		return "Verimatrix VCAS"
	case Invalid:
		return "Invalid"
	default:
		return "Unknown"
	}
}