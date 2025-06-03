package segments

import (
	"regexp"
	"testing"
)

func TestPatEmptyStr(t *testing.T) {
	pattern := ""
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := "In/the/beginning/God/created/the/heavens/and/the/earth"
	want := regexp.MustCompile(`\b` + expected + `\b`)
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q, ""`, input, pattern, output, want)
	}
}

func TestPatEmptyBrackets(t *testing.T) {
	pattern := "[]"
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := ""
	want := regexp.MustCompile(`` + expected + ``)
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, want)
	}
}

func TestPatOneSemicolon(t *testing.T) {
	pattern := "[:]"
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := "In/the/beginning/God/created/the/heavens/and/the/earth"
	want := regexp.MustCompile(`` + expected + ``)
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, want)
	}
}

func TestPatTwoSemicolons(t *testing.T) {
	pattern := "[::]"
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := "In/the/beginning/God/created/the/heavens/and/the/earth"
	want := regexp.MustCompile(`` + expected + ``)
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, want)
	}
}

func TestPatStartZeroStopBlank(t *testing.T) {
	pattern := "[0:]"
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := "In/the/beginning/God/created/the/heavens/and/the/earth"
	want := regexp.MustCompile(`` + expected + ``)
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, want)
	}
}

func TestPatStartZeroStopBlankStepBlank(t *testing.T) {
	pattern := "[0::]"
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := "In/the/beginning/God/created/the/heavens/and/the/earth"
	want := regexp.MustCompile(`` + expected + ``)
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, want)
	}
}

func TestStartSecond(t *testing.T) {
	pattern := "[1:]"
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := "the/beginning/God/created/the/heavens/and/the/earth"
	want := regexp.MustCompile(`` + expected + ``)
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, want)
	}
}

func TestStartThird(t *testing.T) {
	pattern := "[2:]"
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := "beginning/God/created/the/heavens/and/the/earth"
	want := regexp.MustCompile(`` + expected + ``)
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, want)
	}
}

func TestStartFourth(t *testing.T) {
	pattern := "[3:]"
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := "God/created/the/heavens/and/the/earth"
	want := regexp.MustCompile(`` + expected + ``)
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, want)
	}
}

func TestStartOutOfRange(t *testing.T) {
	pattern := "[100:]"
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := ""
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, expected)
	}
}

func TestStopOutOfRange(t *testing.T) {
	pattern := "[0:100]"
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := "In/the/beginning/God/created/the/heavens/and/the/earth"
	want := regexp.MustCompile(`` + expected + ``)
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, want)
	}
}

func TestNegativeStartLast(t *testing.T) {
	pattern := "[-1:]"
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := "earth"
	want := regexp.MustCompile(`` + expected + ``)
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, want)
	}
}

func TestNegativeStartSecondToLast(t *testing.T) {
	pattern := "[-2:]"
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := "the/earth"
	want := regexp.MustCompile(`` + expected + ``)
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, want)
	}
}

func TestNegativeStartOutOfRange(t *testing.T) {
	pattern := "[-100:]"
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := "In/the/beginning/God/created/the/heavens/and/the/earth"
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, expected)
	}
}

func TestNegativeStopAtSecondToLast(t *testing.T) {
	pattern := "[:-2]"
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := "In/the/beginning/God/created/the/heavens/and/the"
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, expected)
	}
}

func TestNegativeStopOutOfRange(t *testing.T) {
	pattern := "[:-100]"
	input := "In/the/beginning/God/created/the/heavens/and/the/earth"
	expected := "In/the/beginning/God/created/the/heavens/and/the/earth"
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, expected)
	}
}

// Negative Step (Reverse Segments)
func TestNegativeStep(t *testing.T) {
	pattern := "[::-1]"
	input := "In/the/beginning"
	expected := "beginning/the/In"
	output := Slice(input, pattern)
	if expected != output {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, expected)
	}
}
