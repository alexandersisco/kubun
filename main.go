package main

import (
	"bufio"
	"fmt"
	"github.com/alexflint/go-arg"
	"io"
	"os"
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of kubun:\n")
	fmt.Fprintf(os.Stderr, "\tkubun '[:-2]' # Select all fields, excluding the last (behavior is like dirname)\n")
	fmt.Fprintf(os.Stderr, "\tkubun '[-1:]' # Select the last field (behavior is like basename)\n")
	fmt.Fprintf(os.Stderr, "\tkubun ',[:]\n' # Select all fields delimited by commas and change the delimiters to newlines\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
}

func main() {
	var path string

	var args struct {
		SlicePat       string `arg:"positional,required" help:"Use Python inspired slice syntax: [start:stop:step]"`
		Input          string `arg:"positional" help:"Input string to select from"`
		ExcludeNewline bool   `arg:"-n,--exclude-newline" help:"Do not output trailing newline"`
	}

	arg.MustParse(&args)

	var slicePat string = args.SlicePat
	var excludeTrailingNewline bool = args.ExcludeNewline

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

	if len(args.Input) > 0 {
		path = args.Input
	}

	newPath := Slice(path, slicePat, "/")
	outputFormat := "%s\n"
	if excludeTrailingNewline {
		outputFormat = "%s"
	}
	fmt.Fprintf(os.Stdout, outputFormat, newPath)
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
