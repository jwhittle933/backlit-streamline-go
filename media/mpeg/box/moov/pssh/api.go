// Package pssh (Protection System Specific Header)
package pssh

import (
	"encoding/hex"
	"fmt"
	"github.com/jwhittle933/streamline/bits/slicereader"
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/fullbox"
)

const (
	PSSH string = "pssh"

	UUIDPlayReady = "9a04f079-9840-4286-ab92-e65be0885f95"
	UUIDWidevine  = "edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"
	UUIDFairPlay  = "94CE86FB-07FF-4F43-ADB8-93D2FA968CA2"
	UUID_VCAS     = "9a27dd82-fde2-4725-8cbc-4234aa06ec09"
)

// Box is ISO BMFF pssh box type
type Box struct {
	fullbox.Box
	SystemID KID
	KIDCount uint32
	KIDs     []KID
	DataSize uint32
	Data     []byte
}

type KID [16]byte

func (k KID) String() string {
	h := hex.EncodeToString(k[:])
	return fmt.Sprintf("%s-%s-%s-%s-%s", h[:8], h[8:12], h[12:16], h[16:20], h[20:32])
}

func New(i *box.Info) box.Boxed {
	return &Box{
		*fullbox.New(i),
		[16]byte{},
		0,
		make([]KID, 0),
		0,
		make([]byte, 0),
	}
}

func (Box) Type() string {
	return PSSH
}

func (b Box) String() string {
	return fmt.Sprintf(
		"%s sys_id=%s, kids=%d, data_size=%d",
		b.Info(),
		b.SystemID,
		b.KIDCount,
		b.DataSize,
	)
}

// Write satisfies the io.Writer interface
func (b *Box) Write(src []byte) (int, error) {
	sr := slicereader.New(src)
	b.WriteVersionAndFlags(sr)

	copy(b.SystemID[:], sr.Slice(16))

	if b.Version > 0 {
		b.KIDCount = sr.Uint32()
		for i := uint32(0); i < b.KIDCount; i++ {
			var kid KID
			copy(kid[:], sr.Slice(16))
			b.KIDs = append(b.KIDs, kid)
		}
	}

	b.DataSize = sr.Uint32()
	b.Data = make([]byte, b.DataSize)
	sr.Read(b.Data)

	return len(src), nil
}
func (b *Box) Size() uint64 {
	size := uint64(12 + 16 + 4 + len(b.Data))
	if b.Version > 0 {
		size += uint64(4 + 16*len(b.KIDs))
	}

	return size
}

func systemIDString(id KID) string {
	switch id.String() {
	default:
		return "Unknown"
	case UUIDPlayReady:
		return "PlayReady"
	case UUIDWidevine:
		return "Widevine"
	case UUIDFairPlay:
		return "FairPlay"
	case UUID_VCAS:
		return "Verimatrix VCAS"
	}
}
