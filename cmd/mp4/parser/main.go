package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/jwhittle933/streamline/media/mp4"
	"github.com/jwhittle933/streamline/result"
)

func main() {
	inputFile := flag.String("file", "", "File to parse")
	dump := flag.Bool("dump", false, "Dump the hex values")
	flag.Parse()
	if inputFile == nil || *inputFile == "" {
		exitOnError(fmt.Errorf("no file provided. Exiting"), 0)
	}

	result.Pipe(
		mp4.Open(*inputFile),
		exitOnDump(*dump),
		read(),
		printMP4(),
		//JSON,
	)
}

func JSON(data interface{}) result.Result {
	b, _ := json.MarshalIndent(data.(*mp4.MP4), "", "  ")
	fmt.Println(string(b))

	return result.Wrap(data)
}

func read() result.Binder {
	return func(data interface{}) result.Result {
		if err := data.(*mp4.MP4).ReadAll(); err != nil {
			return result.WrapErr(err)
		}

		return result.Wrap(data)
	}
}

func exitOnDump(shouldDump bool) result.Binder {
	return func(data interface{}) result.Result {
		if shouldDump {
			fmt.Println(data.(*mp4.MP4).Hex())
			os.Exit(0)
		}

		return result.Wrap(data)
	}
}

func printMP4() result.Binder {
	return func(data interface{}) result.Result {
		m := data.(*mp4.MP4)
		fmt.Printf("[\033[1;35mmp4\033[0m] size=%d, fragmented=%+v, boxes=%d\n", m.Size, m.IsFragmented(), len(m.Children))

		for _, c := range m.Children {
			fmt.Println(c)
		}
		return result.Wrap(data)
	}
}

func exitOnError(err error, code int) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(code)
	}
}
