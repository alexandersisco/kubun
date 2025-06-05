package segments

import (
	"strconv"
	"strings"
)

func Slice(s string, pat string, delimiter string) string {
	s = strings.TrimRight(s, "\r\n")
	if pat == "" || pat == "[:]" || pat == "[::]" || pat == "[0:]" || pat == "[0::]" {
		return s
	}
	if pat == "[]" {
		return ""
	}
	s = strings.Trim(s, "\n\r ")

	d := ExtractDelimiters(pat)
	if d.delimiter != "" {
		delimiter = d.delimiter
	}
	segments := strings.Split(s, delimiter)
	start, stop, step := ParseSlicePattern(len(segments), d.slicePattern)
	sliced := segments[start:stop]

	if step < 0 {
		ReverseSegments(sliced)
	}
	joinDelimiter := delimiter
	if d.shouldReplaceDelimiter {
		joinDelimiter = d.newDelimiter
	}
	s = strings.Join(sliced, joinDelimiter)
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
