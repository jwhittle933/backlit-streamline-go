package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jwhittle933/streamline/media/mp4"
)

func main() {
	inputFile := flag.String("file", "", "File to parse")
	dump := flag.Bool("dump", false, "Dump the hex values")
	flag.Parse()
	if inputFile == nil || *inputFile == "" {
		exitOnError(fmt.Errorf("no file provided. Exiting"), 0)
	}

	p, err := filepath.Abs(*inputFile)
	exitOnError(err, 0)

	if ext := filepath.Ext(p); ext == "" {
		exitOnError(
			fmt.Errorf("Invalid file ext \"%s\"\nExiting...\n", *inputFile),
			0,
		)
	}

	file, err := os.Open(p)
	exitOnError(err, 1)
	defer file.Close()

	m, err := mp4.New(file)
	exitOnError(err, 1)

	if *dump {
		fmt.Println(m.Hex())
		return
	}

	exitOnError(m.ReadAll(), 1)
	fmt.Printf("[\033[1;35mmp4\033[0m] size=%d, boxes=%d\n", m.Size, len(m.Boxes))
	for _, b := range m.Boxes {
		fmt.Println(b)
	}
}

func exitOnError(err error, code int) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(code)
	}
}
