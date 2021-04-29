package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/jwhittle933/streamline/pkg/media/mp4"
	"io/ioutil"
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

	fmt.Println("File Name:", file.Name())
	info, err := file.Stat()
	if err != nil {
		fmt.Printf("Could read file info: %s\nExiting...\n", err.Error())
		os.Exit(0)
	}

	fmt.Printf("File Size: %dMB\n", info.Size()/1024)

	if *dump {
		src, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Printf("Could read file: %s\nExiting...\n", err.Error())
			os.Exit(0)
		}

		fmt.Print(hex.Dump(src))
		return
	}

	m, err := mp4.New(file)
	exitOnError(err, 1)

	fmt.Printf("\n")

	ftyp, err := m.ReadInfo()
	exitOnError(err, 1)
	fmt.Println("Offset:", ftyp.Offset)
	fmt.Println("Type (raw):", ftyp.Type[:])
	fmt.Println("Type:", string(ftyp.Type[:]))
	fmt.Println("Size:", ftyp.Size)

	fmt.Printf("\n")

	bi2, err := m.ReadInfo()
	exitOnError(err, 1)
	fmt.Println("Offset:", bi2.Offset)
	fmt.Println("Type (raw):", bi2.Type[:])
	fmt.Println("Type:", string(bi2.Type[:]))
	fmt.Println("Size:", bi2.Size)

	bi3, err := m.ReadInfo()
	exitOnError(err, 1)
	fmt.Println("Offset:", bi3.Offset)
	fmt.Println("Type (raw):", bi3.Type[:])
	fmt.Println("Type:", string(bi3.Type[:]))
	fmt.Println("Size:", bi3.Size)
}

func exitOnError(err error, code int) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(code)
	}
}
