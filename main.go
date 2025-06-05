package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/alexandersisco/kubun/segments"
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of kubun:\n")
	fmt.Fprintf(os.Stderr, "\tkubun -s '[:-2]' # Select all segments, excluding the last (behavior is like dirname)\n")
	fmt.Fprintf(os.Stderr, "\tkubun -s '[-1:]' # Select the last segment (behavior is like basename)\n")
	fmt.Fprintf(os.Stderr, "\tkubun -d ',' -s '[1:]' # Select all segments divided by ',', skipping the first one\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	var path string

	flag.Usage = Usage

	var slicePat string
	flag.StringVar(&slicePat, "s", "[::]", "select segments from the string based on the slice pattern: [start:stop:step]")

	var delimiter string = "/"
	flag.StringVar(&delimiter, "d", "/", "the delimiter that divides the string into segments")

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
	fmt.Fprint(os.Stdout, newPath)
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
