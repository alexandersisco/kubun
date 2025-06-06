package main

import (
	"fmt"
	"testing"
)

type testCaseDelimters struct {
	input                  string
	slicePattern           string
	delimiter              string
	newDelimiter           string
	shouldReplaceDelimiter bool
}

func TestDelimiters(t *testing.T) {
	testCases := []testCaseDelimters{
		{
			"/[:]\\",
			"[:]",
			"/",
			"\\",
			true,
		},
		{
			"/[:]",
			"[:]",
			"/",
			"",
			false,
		},
		{
			"[:]",
			"[:]",
			"",
			"",
			false,
		},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		d := ExtractDelimiters(test.input)
		if d.delimiter != test.delimiter || d.newDelimiter != test.newDelimiter || d.slicePattern != test.slicePattern || d.shouldReplaceDelimiter != test.shouldReplaceDelimiter {
			failCount++
			t.Errorf(`
----------------------------------
FAILED: Extracting delimiters from %v failed to produce the expected values.

Expected values:
----------------
SlicePat:                 %v
Existing delimiter:       %v
Replacement delimiter:    %v
Should replace:           %v

Actual values:
----------------
SlicePat:                 %v
Existing delimiter:       %v
Replacement delimiter:    %v
Should replace:           %v
`, test.input, test.slicePattern, test.delimiter, test.newDelimiter, test.shouldReplaceDelimiter, d.slicePattern, d.delimiter, d.newDelimiter, d.shouldReplaceDelimiter)

		} else {
			passCount++
			fmt.Printf(`----------------------------------
PASS: Extracted the following values from %v --> delA = "%v", delB = "%v", slicePat = "%v", replace = %v
`, test.input, test.delimiter, test.newDelimiter, test.slicePattern, test.shouldReplaceDelimiter)
		}
	}
	fmt.Println("----------------------------------")

	fmt.Println("")
	fmt.Println("")
	fmt.Printf(`----------------------------------
Failed: %v
Passed: %v
`, failCount, passCount)
	fmt.Println("----------------------------------")
	fmt.Println("")
}
