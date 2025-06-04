package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/alexandersisco/pathy/segments"
)

func main() {
	var path string

	var slicePat string
	flag.StringVar(&slicePat, "slice", "[::]", "a pattern for taking slices from the path")

	var delimiter string = "/"
	flag.StringVar(&delimiter, "delimiter", "/", "character or string that divides the segments")

	flag.Parse()

	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if (fi.Mode() & os.ModeCharDevice) != 0 {
		// Stdin is connected to a terminal
	} else if (fi.Mode() & os.ModeNamedPipe) != 0 {
		// Stdin is connected to a pipe
		path, _ = ReadStdIn()
	} else {
		// Stdin is a redirected file
		path, _ = ReadStdIn()
	}

	newPath := segments.Slice(path, slicePat, delimiter)
	println(newPath)
}

func ReadStdIn() (string, error) {
	rdr := bufio.NewReader(os.Stdin)

	switch line, err := rdr.ReadString('\n'); err {
	case nil:
		return line, nil
	case io.EOF:
		return "", err
	default:
		fmt.Fprintln(os.Stderr, "error:", err)
		return "", err
	}
}
