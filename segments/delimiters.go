package segments

import "strings"

type ExtractedDelimiters struct {
	slicePattern           string
	delimiter              string
	newDelimiter           string
	shouldReplaceDelimiter bool
}

func ExtractDelimiters(pat string) ExtractedDelimiters {
	before, after, _ := strings.Cut(pat, "[")
	pat, after, _ = strings.Cut(after, "]")

	return ExtractedDelimiters{
		delimiter:              strings.Replace(before, "\\n", "\n", -1),
		newDelimiter:           strings.Replace(after, "\\n", "\n", -1),
		slicePattern:           "[" + pat + "]",
		shouldReplaceDelimiter: before != "" && after != "",
	}
}
