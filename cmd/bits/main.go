package main

import (
	"fmt"

	"github.com/jwhittle933/streamline/bits"
)

func main() {
	br := bits.NewReader([]byte{0xFF, 0x00, 0x0F})

	bitVal, err := br.ReadBits(8)
	if err != nil {
		fmt.Println(err.Error())
	}

	bitString := ""
	for _, val := range bitVal {
		bitString += fmt.Sprintf("%s", val)
	}

	fmt.Println(bitString)

	num := uint32(111151000)
	fmt.Printf("\nNum: %b\n", num)
	fmt.Printf("Num >> 24: %b, %d\n", num >> 24, num >> 24)
}

