package segments

import (
	"strconv"
	"strings"
)

func Slice(s string, pat string, delimiter string) string {
	if pat == "" || pat == "[:]" || pat == "[::]" || pat == "[0:]" || pat == "[0::]" {
		return s
	}
	if pat == "[]" {
		return ""
	}
	s = strings.Trim(s, "\n\r ")
	segments := strings.Split(s, delimiter)
	start, stop, step := ParseSlicePattern(len(segments), pat)
	sliced := segments[start:stop]

	if step < 0 {
		ReverseSegments(sliced)
	}
	s = strings.Join(sliced, delimiter)
	return s
}

func ReverseSegments(segments []string) {
	for i, j := 0, len(segments)-1; i < j; i, j = i+1, j-1 {
		segments[i], segments[j] = segments[j], segments[i]
	}
}

func ParseSlicePattern(segmentCount int, pattern string) (int, int, int) {
	pattern = pattern[1 : len(pattern)-1]

	parts := strings.Split(pattern, ":")
	if len(parts) < 2 {
		return 0, segmentCount, 1
	}

	// Start
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		start = 0
	}

	// Negative indexing for start
	if start < 0 {
		start = max(segmentCount+start, 0)
	}

	if start > segmentCount {
		start = segmentCount
	}

	// Stop
	stop, err := strconv.Atoi(parts[1])
	if err != nil || stop > segmentCount {
		stop = segmentCount
	}

	// Negative indexing for stop
	if stop < 0 {
		stop = segmentCount + stop + 1
		if stop < 0 {
			stop = segmentCount
		}
	}

	// Step
	if len(parts) < 3 {
		return start, stop, 1
	}
	step, err := strconv.Atoi(parts[2])
	if err != nil {
		step = 1
	}

	return start, stop, step
}
