// Package boxtype
// see https://github.com/abema/go-mp4/blob/8afbbcf4788ca39d4147047bce9dc915b013c9fa/box_types.go#L253
package boxtype

import "fmt"

type BoxType [4]byte

func New(code [4]byte) BoxType {
	return BoxType{code[0], code[1], code[2], code[3]}
}

func (b BoxType) String() string {
	return string(b[:])
}

func (b BoxType) HexString() string {
	return fmt.Sprintf("0x%02x%02x%02x%02x", b[0], b[1], b[2], b[3])
}
func isASCII(c byte) bool {
	return c >= 0x20 && c <= 0x7e
}

func isPrintable(c byte) bool {
	return isASCII(c) || c == 0xa9
}
