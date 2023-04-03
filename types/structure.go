package types

import (
	"sort"
	"strconv"
	"strings"

	"github.com/json-go/format"
)

// Convert a JSON object to a Go structure type. The object is represented as a
// map with the field names and the values, which are used to determine the
// appropriate Go data types to assign to the associated struct field values.
func structure(actual map[string]interface{}, depth int, baseType string, camel bool, typeMap map[string]string) string {
	names := []string{}
	types := []string{}
	tags := []string{}

	keys := []string{}
	for key := range actual {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		value := actual[key]
		names = append(names, format.PascalCase(key, camel))
		types = append(types, Element(value, depth+1, baseType, camel, typeMap))
		tags = append(tags, format.Tag(key))
	}

	maxName := 0
	maxType := 0

	for i := 0; i < len(names); i++ {
		if w := format.Width(names[i]); w > maxName {
			maxName = w
		}

		if w := format.Width(types[i]); w > maxType {
			maxType = w
		}
	}

	result := strings.Builder{}

	result.WriteString(format.Pad("", depth*2))
	result.WriteString("struct {\n")

	for i := 0; i < len(names); i++ {
		result.WriteString("  ")
		result.WriteString(format.Pad(names[i], maxName))
		result.WriteString(" ")
		result.WriteString(format.Indent(types[i], 0, maxType))
		result.WriteString(" ")
		result.WriteString(tags[i])
		result.WriteString("\n")
	}

	result.WriteString("}\n")

	t := result.String()
	if typeName, found := typeMap[t]; found {
		return typeName
	}

	if depth == 0 {
		return t
	}

	if baseType == "" {
		baseType = "generated"
	}

	typeName := baseType + "Type" + strconv.Itoa(1+len(typeMap))
	typeMap[t] = typeName

	return typeName
}
