package main

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/json-go/types"
)

// Convert converts a json payload into a Go definition string.  The
// base type name is provided, which becomes the name of the top-level
// type in the resulting Go code. The camel flag is used to indicate if
// field names are coerced to camel case in the Go code (which determines
// if they are visible outside the package). Normally camel case should be
// off to prevent the resulting data types from containing only private
// data items which prevents JSON from working correctly.
func Convert(data []byte, typeName string, camel bool) (string, error) {
	var x interface{}

	typeMap := map[string]string{}

	err := json.Unmarshal(data, &x)
	if err != nil {
		return "", err
	}

	dataType := types.Element(x, 0, typeName, camel, typeMap)
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
