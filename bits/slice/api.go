package slice

import "errors"

// Uint8 consumes the first byte of `data`
func Uint8(data []byte) ([]byte, uint8, error) {
	if len(data) == 0 {
		return nil, 0, errors.New("attempted read from empty slice")
	}

	return data[1:], data[0], nil
}
