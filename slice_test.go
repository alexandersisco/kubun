package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	name      string
	delimiter string
	pattern   string
	input     string
	expected  string
}

func Test(t *testing.T) {
	testCases := []testCase{
		{
			"Select everything without a slice pattern",
			"/",
			"",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
		},
		{
			"Select nothing",
			"/",
			"[]",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
			"",
		},
		{
			"Select every segment with [:]",
			"/",
			"[:]",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
		},
		{
			"Select every segment with [::]",
			"/",
			"[::]",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
		},
		{
			"Select every segment with [0:]",
			"/",
			"[0:]",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
		},
		{
			"Select every segment with [0::]",
			"/",
			"[0::]",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
		},
		{
			"Skipt the first segment and select everything else",
			"/",
			"[1:]",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
			"the/beginning/God/created/the/heavens/and/the/earth",
		},
		{
			"Skip the first two segments and select everything else",
			"/",
			"[2:]",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
			"beginning/God/created/the/heavens/and/the/earth",
		},
		{
			"Start out of range",
			"/",
			"[100:]",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
			"",
		},
		{
			"Select last segment with negative start index (like basename command)",
			"/",
			"[-1:]",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
			"earth",
		},
		{
			"Select the last two segments with negative start index",
			"/",
			"[-2:]",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
			"the/earth",
		},
		{
			"Negative start out of range",
			"/",
			"[-100:]",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
		},
		{
			"Negative start out of range",
			"/",
			"[-100:]",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
		},
		{
			"Select everything but the last segment with [:-2] (like dirname command)",
			"/",
			"[:-2]",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
			"In/the/beginning/God/created/the/heavens/and/the",
		},
		{
			"Negative stop out of range",
			"/",
			"[:-100]",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
			"In/the/beginning/God/created/the/heavens/and/the/earth",
		},
		{
			"Reverse segments with negative step",
			"/",
			"[::-1]",
			"the/heavens/and/the/earth",
			"earth/the/and/heavens/the",
		},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := Slice(test.input, test.pattern, test.delimiter)
		if output != test.expected {
			failCount++
			t.Errorf(`
----------------------------------
FAILED:    %v
Delimter:  %v
Slice:     %v
Input:     %v
Expected:  %v
Actual:    %v
`, test.name, test.delimiter, test.pattern, test.input, test.expected, output)

		} else {
			passCount++
			fmt.Printf(`----------------------------------
PASS: %v
`, test.name)
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
