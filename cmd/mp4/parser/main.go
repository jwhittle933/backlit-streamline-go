package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	inputFile := flag.String("file", "", "File to parse")
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

	fmt.Printf("File Size: %dMB\n", info.Size() / 1024)
	src, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Could read file: %s\nExiting...\n", err.Error())
		os.Exit(0)
	}

	fmt.Print(hex.Dump(src))
}


