package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jwhittle933/streamline/media/mpeg"
	"github.com/jwhittle933/streamline/result"
)

func main() {
	inputFile := flag.String("file", "", "File to parse")
	dump := flag.Bool("dump", false, "Dump the hex values")
	flag.Parse()
	if inputFile == nil || *inputFile == "" {
		exitOnError(fmt.Errorf("no file provided. Exiting"), 0)
	}

	p := result.Pipe(
		mpeg.Open(*inputFile),
		exitOnDump(*dump),
		read(),
		printMP4(),
	)

	if p.IsErr() {
		fmt.Println("Error: ", p.Err().Error())
	}
}

func read() result.Binder {
	return func(data interface{}) result.Result {
		if err := data.(*mpeg.MPEG).ReadAll(); err != nil {
			return result.WrapErr(err)
		}

		return result.Wrap(data)
	}
}

func exitOnDump(shouldDump bool) result.Binder {
	return func(data interface{}) result.Result {
		if shouldDump {
			fmt.Println(data.(*mpeg.MPEG).Hex())
			os.Exit(0)
		}

		return result.Wrap(data)
	}
}

func printMP4() result.Binder {
	return func(data interface{}) result.Result {
		m := data.(*mpeg.MPEG)
		fmt.Printf("%s\n", m)
		return result.Wrap(data)
	}
}

func exitOnError(err error, code int) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(code)
	}
}

func sized(s int) string {
	var out string

	for s > 1024 {
		//
	}

	return out
}