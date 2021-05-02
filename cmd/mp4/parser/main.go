package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jwhittle933/streamline/pkg/media/mp4"
)

func main() {
	inputFile := flag.String("file", "", "File to parse")
	dump := flag.Bool("dump", false, "Dump the hex values")
	flag.Parse()
	if inputFile == nil || *inputFile == "" {
		fmt.Println("No file provided Exiting...")
		os.Exit(0)
	}

	p, err := filepath.Abs(*inputFile)
	if err != nil {
		fmt.Printf("Invalid filepath \"%s\"\nExiting...\n", *inputFile)
		os.Exit(0)
	}

	if ext := filepath.Ext(p); ext == "" {
		fmt.Printf("Invalid file ext \"%s\"\nExiting...\n", *inputFile)
		os.Exit(0)
	}

	fmt.Println("File: ", p)
	file, err := os.Open(p)
	defer file.Close()
	if err != nil {
		fmt.Printf("Could not open file \"%s\": %s\nExiting...\n", p, err.Error())
		os.Exit(1)
	}

	m, err := mp4.New(file)
	exitOnError(err, 1)

	if *dump {
		fmt.Println(m.Hex())
		return
	}

	fmt.Printf("\n")

	exitOnError(m.ReadAll(), 1)
	fmt.Printf("[mp4] size=%d\n", m.Size)
	for _, b := range m.Boxes {
		fmt.Println(b.String())
	}
}

func exitOnError(err error, code int) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(code)
	}
}
