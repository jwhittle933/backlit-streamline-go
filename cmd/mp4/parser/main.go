package main

import (
	"flag"
	"fmt"
	mp42 "github.com/jwhittle933/streamline/media/mp4"
	"os"
	"path/filepath"
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

	m, err := mp42.New(file)
	exitOnError(err, 1)

	if *dump {
		fmt.Println(m.Hex())
		return
	}

	fmt.Printf("\n")

	exitOnError(m.ReadAll(), 1)
	fmt.Printf("[\033[1;35mmp4\033[0m] size=%d, boxes=%d\n", m.Size, len(m.Boxes))
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
