package main

import "fmt"
import "bufio"
import "io"
import "os"
import "strings"

func main() {
	path := "~/workspace/goproj/pathy/main.go"

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

	// ListPathSegments(path)

	dir_path := DirPath(path)
	fmt.Println(dir_path)
}

func DirPath(path string) string {
	segments := strings.Split(path, "/")

	length := len(segments)

	str := ""
	for i, v := range segments {
		if i < length-1 {
			str += v + "/"
		}
	}

	return str[0 : len(str)-1]
}

func ListPathSegments(path string) {
	segments := strings.Split(path, "/")

	for i, v := range segments {
		fmt.Println(i, v)
	}
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
