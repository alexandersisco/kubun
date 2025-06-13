package main

import (
	"bufio"
	"fmt"
	"github.com/alexflint/go-arg"
	"io"
	"os"
)

type args struct {
	SlicePat       string `arg:"positional,required" help:"Use Python inspired slice syntax: [start:stop:step]"`
	Input          string `arg:"positional" help:"Input string to select from"`
	ExcludeNewline bool   `arg:"-n,--exclude-newline" help:"Do not output trailing newline"`
}

func (args) Description() string {
	return `
Python-style slicing on the command line for delimiter separated text

Examples:
  kubun '[:]' /usr/bin/sort               -> /usr/bin/sort
  kubun '[-1:]' /usr/bin/sort             -> sort
  kubun '[-2:]' /usr/bin/sort             -> bin/sort

  Replacing delimiters:
  kubun '/[:]\'                           -> \usr\bin\sort
  kubun '/[1:], '                         -> usr, bin, sort
  kubun '/[1:]\n'                         -> usr
                                             bin
                                             sort

  Reverse fields:
  kubun '[::-1]' /usr/bin/sort            -> sort/bin/usr/

  Stdin:
  echo "/usr/bin/sort" | kubun '[-3:]'    -> usr/bin/sort
`
}

func main() {
	var path string

	var args args
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
