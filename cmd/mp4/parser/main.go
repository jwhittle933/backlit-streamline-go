package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
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

	fmt.Printf("\n")

	bi, err := readBoxInfo(file)
	exitOnError(err, 1)
	fmt.Println("Offset:", bi.Offset)
	fmt.Println("Type:", string(bi.Type[:]))
	fmt.Println("Size:", bi.Size)

	fmt.Printf("\n")

	bi2, err := readBoxInfo(file)
	exitOnError(err, 1)
	fmt.Println("Offset:", bi.Offset)
	fmt.Println("Type:", string(bi2.Type[:]))
	fmt.Println("Size:", bi2.Size)
}

func readBoxInfo(r io.ReadSeeker) (*box.Info, error) {
	off, _ := r.Seek(0, io.SeekCurrent)

	bi := &box.Info{
		Offset:     uint64(off),
		HeaderSize: box.SmallHeader,
	}

	buf := bytes.NewBuffer(make([]byte, 0, bi.HeaderSize))
	if _, err := io.CopyN(buf, r, box.SmallHeader); err != nil {
		return nil, err
	}

	data := buf.Bytes()
	bi.Size = uint64(binary.BigEndian.Uint32(data))
	bi.Type = [4]byte{data[4], data[5], data[6], data[7]}

	if bi.Size == 0 {
		off, _ = r.Seek(0, io.SeekEnd)
		bi.Size = uint64(off) - bi.Offset
		bi.ExtendToEOF = true
		if _, err := bi.SeekPayload(r); err != nil {
			return nil, err
		}

		return bi, nil
	}

	if bi.Size == 1 {
		buf.Reset()
		if _, err := io.CopyN(buf, r, box.LargeHeader-box.SmallHeader); err != nil {
			return nil, err
		}

		bi.HeaderSize += box.LargeHeader - box.SmallHeader
		bi.Size = binary.BigEndian.Uint64(buf.Bytes())
		return bi, nil
	}

	return bi, nil
}

func exitOnError(err error, code int) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(code)
	}
}
