package segments

import (
	"strconv"
	"strings"
)

func Slice(s string, pat string) string {
	if pat == "" || pat == "[:]" || pat == "[::]" || pat == "[0:]" || pat == "[0::]" {
		return s
	}
	if pat == "[]" {
		return ""
	}
	start, stop := ParseSlicePattern(s, pat)
	segments := strings.Split(s, "/")
	sliced := segments[start:stop]
	s = strings.Join(sliced, "/")
	return s
}

func ParseSlicePattern(s string, pattern string) (int, int) {
	pattern = pattern[1 : len(pattern)-1]
	segments := strings.Split(s, "/")

	parts := strings.Split(pattern, ":")
	if len(parts) < 2 {
		return 0, len(segments)
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		start = 0
	}

	stop, err := strconv.Atoi(parts[1])
	if err != nil {
		stop = len(segments)
	}
	return start, stop
}
