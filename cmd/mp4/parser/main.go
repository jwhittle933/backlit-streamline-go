package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
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

	//buf := make([]byte, 8)
	//_, err = file.Read(buf)
	//exitOnError(err, 1)
	//fmt.Println(buf)
	//fmt.Println(string(buf))
	//fmt.Println(hex.Dump(buf))
	//
	//fmt.Println("size:", buf[3])
	//fmt.Println("type:", string(buf[4:]))
	//dataBuf := make([]byte, buf[3])
	//_, err = file.Read(dataBuf)
	//exitOnError(err, 1)
	//fmt.Println(dataBuf)
	//fmt.Println(string(dataBuf))
	//fmt.Println(hex.Dump(dataBuf))

	fmt.Printf("\n")

	offset, name, data, _, err := chunk(file)
	exitOnError(err, 1)
	fmt.Println("offset:", offset)
	fmt.Println("name:", name)
	fmt.Println("data:", data)
	fmt.Println("Buffer length:", len(append(data, name...)))

	fmt.Printf("\n")

	offset, name, data, _, err = chunk(file)
	exitOnError(err, 1)
	fmt.Println("offset:", offset)
	fmt.Println("name:", name)
	fmt.Println("data:", data)
	fmt.Println("Buffer length:", len(append(data, name...)))
}

func chunk(r io.ReadSeeker) (int, string, []byte, bool, error) {
	sizeAndName := buffer(8)
	if _, err := r.Read(sizeAndName); err != nil {
		return 0, "", nil, true, err
	}
	fmt.Println(sizeAndName)

	size := int(sizeAndName[3])
	data := buffer(size - len(sizeAndName))
	if _, err := r.Read(data); err != nil {
		return size, "", nil, true, err
	}

	name := string(sizeAndName[4:])
	return size, name, data[4:], false, nil
}

func buffer(size int) []byte {
	return make([]byte, size)
}

func exitOnError(err error, code int) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(code)
	}
}
