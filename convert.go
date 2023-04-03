package main

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func Convert(data []byte, typeName string, camel bool) (string, error) {
	var x interface{}

	typeMap := map[string]string{}

	err := json.Unmarshal(data, &x)
	if err != nil {
		return "", err
	}

	dataType := convert(x, 0, typeName, camel, typeMap)
	if typeName != "" {
		dataType = "type " + typeName + " " + dataType
	}

	if len(typeMap) == 0 {
		return dataType, nil
	}

	result := strings.Builder{}

	invertedMap := map[string]string{}
	for key, value := range typeMap {
		invertedMap[value] = key
	}

	keys := []string{}
	for key := range invertedMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		typeDef := invertedMap[key]

		result.WriteString("\n")
		result.WriteString("type ")
		result.WriteString(key)
		result.WriteString(" ")
		result.WriteString(strings.TrimSpace(typeDef))
		result.WriteString("\n")
	}

	result.WriteString("\n")
	result.WriteString(dataType)

	return result.String(), nil
}

func convert(x interface{}, depth int, baseType string, camel bool, typeMap map[string]string) string {
	switch actual := x.(type) {
	case map[string]interface{}:
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
			names = append(names, pascalCase(key, camel))
			types = append(types, convert(value, depth+1, baseType, camel, typeMap))
			tags = append(tags, tag(key))
		}

		maxName := 0
		maxType := 0

		for i := 0; i < len(names); i++ {
			if w := width(names[i]); w > maxName {
				maxName = w
			}

			if w := width(types[i]); w > maxType {
				maxType = w
			}
		}

		result := strings.Builder{}

		result.WriteString(pad("", depth*2))
		result.WriteString("struct {\n")

		for i := 0; i < len(names); i++ {
			result.WriteString("  ")
			result.WriteString(pad(names[i], maxName))
			result.WriteString(" ")
			result.WriteString(indent(types[i], 0, maxType))
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

	case bool:
		return "bool"

	case int:
		return "int"

	case int64:
		return "int64"

	case string:
		return "string"

	case float64:
		if actual == math.Floor(actual) {
			return "int"
		}

		return "float64"

	case []interface{}:
		var t string

		for index, value := range actual {
			if index == 0 {
				t = convert(value, depth+1, baseType, camel, typeMap)
			} else {
				if t != convert(value, depth+1, baseType, camel, typeMap) {
					return "[]interface{}"
				}
			}
		}

		return "[]" + t

	default:
		return "<unknown>"
	}
}

func pascalCase(s string, camel bool) string {
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

func tag(key string) string {
	return fmt.Sprintf("`json:\"%s,omitempty\"`", key)
}

func pad(s string, size int) string {
	for len(s) < size {
		s = s + " "
	}

	return s
}

func width(s string) int {
	max := 0
	lines := strings.Split(s, "\n")

	for _, value := range lines {
		if max < len(value) {
			max = len(value)
		}
	}

	return max
}

func indent(s string, spacing, width int) string {
	lines := strings.Split(s, "\n")
	result := strings.Builder{}

	for index, value := range lines {
		if value == "\n" {
			continue
		}

		if index == len(lines)-1 {
			value = strings.TrimSuffix(value, "\n")
		}

		result.WriteString(pad("", spacing))
		result.WriteString(pad(value, width))

		if index < len(lines)-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}
