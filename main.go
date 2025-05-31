package main

import "fmt"
import "strings"

func main() {
	path := "~/workspace/goproj/pathy/main.go"

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
