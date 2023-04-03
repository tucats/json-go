package format

import "strings"

// Pad will add spaces to the end of a string value until the
// resulting string is the given length. If a string is passed
// that is already larger than the length, it is returned as-is.
func Pad(s string, size int) string {
	for len(s) < size {
		s = s + " "
	}

	return s
}

// Width determines the widest part of a string. If the string has
// no newline characters, the width is the same as the len() of the
// string. If there are newlines, the width is the len() of the
// longest line in the string.
func Width(s string) int {
	max := 0
	lines := strings.Split(s, "\n")

	for _, value := range lines {
		if max < len(value) {
			max = len(value)
		}
	}

	return max
}

// Indent and pad a string value. The spacing parameter
// indicates how many blanks to add before each line of the
// output, and the width indicates how wide each line must
// be. The given string is broken into a list of strings
// on the newline character, and each line is output with the
// given spacing and width values.
func Indent(s string, spacing, width int) string {
	lines := strings.Split(s, "\n")
	result := strings.Builder{}

	for index, value := range lines {
		if value == "\n" {
			continue
		}

		if index == len(lines)-1 {
			value = strings.TrimSuffix(value, "\n")
		}

		result.WriteString(Pad("", spacing))
		result.WriteString(Pad(value, width))

		if index < len(lines)-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}
