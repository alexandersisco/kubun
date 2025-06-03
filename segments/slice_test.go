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
	if !want.MatchString(output) {
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
	if !want.MatchString(output) {
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
	if !want.MatchString(output) {
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
	if !want.MatchString(output) {
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
	if !want.MatchString(output) {
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
	if !want.MatchString(output) {
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
	if !want.MatchString(output) {
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
	if !want.MatchString(output) {
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
	if !want.MatchString(output) {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n\n", output)
		t.Errorf(`Slice("%s", "%s") = %q, want match for %#q`, input, pattern, output, want)
	}
}
