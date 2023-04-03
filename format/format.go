package format

import (
	"fmt"
	"strings"
	"unicode"
)

// PascalCase applies pascal or camel case on a string. Pascal case
// specifies that the first character must be an uppercase character;
// camel case forces the first character to be a lowercase character.
func PascalCase(s string, camel bool) string {
	result := strings.Builder{}

	for n, ch := range s {
		if n == 0 {
			if camel {
				ch = unicode.ToLower(ch)
			} else {
				ch = unicode.ToUpper(ch)
			}
		}

		result.WriteRune(ch)
	}

	return result.String()
}

// Tag generates the Go json tag string for a given field name.
func Tag(key string) string {
	return fmt.Sprintf("`json:\"%s,omitempty\"`", key)
}
